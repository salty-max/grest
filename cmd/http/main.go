package main

import (
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/salty-max/grest/config"
	mw "github.com/salty-max/grest/internal/middlewares"
	"github.com/salty-max/grest/internal/routes"
	"github.com/salty-max/grest/internal/storage"
	"github.com/salty-max/grest/pkg/shutdown"
)

func main() {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
	})

	// setup exit code for graceful shutdown
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	// load config
	env, err := config.LoadConfig()
	if err != nil {
		logger.Error(err)
		exitCode = 1
		return
	}

	// run the server
	cleanup, err := run(env)

	// run the cleanup after the server is terminated
	defer cleanup()
	if err != nil {
		logger.Error(err)
		exitCode = 1
		return
	}

	// ensure the server is shutdown gracefully & app runs
	shutdown.Gracefully()
}

func run(env config.EnvVars) (func(), error) {
	app, cleanup, err := buildServer(env)
	if err != nil {
		return nil, err
	}

	// start the server
	go func() {
		app.Listen("0.0.0.0:" + env.PORT)
	}()

	// return a function to close the server and database
	return func() {
		cleanup()
		app.Shutdown()
	}, nil
}

func buildServer(env config.EnvVars) (*fiber.App, func(), error) {
	log.Info("connecting to database...")
	// init the storage
	db, err := storage.Connect(env)
	if err != nil {
		return nil, nil, err
	}
	log.Info("successfully connected to database!")

	// create the fiber app
	app := fiber.New()

	// add middleware
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(mw.JSONMiddleWare)

	timeout_delay, err := strconv.ParseInt(os.Getenv("TIMEOUT_DELAY"), 10, 64)
	if err != nil {
		log.Error(err)
	}
	app.Use(mw.TimeoutMiddleware(time.Duration(timeout_delay) * time.Second))

	routes.SetupRoutes(app, db)

	return app, func() {
		storage.Close(db)
	}, nil
}
