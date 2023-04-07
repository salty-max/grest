package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	ct "github.com/salty-max/grest/internal/controllers"
	"github.com/salty-max/grest/internal/storage"
)

func SetupRoutes(app *fiber.App, db *storage.Database) {
	// add root route
	app.Get("/", func(c *fiber.Ctx) error {
		time.Sleep(20 * time.Second)
		return c.SendString("Jaffa, kree!")
	})

	// add health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Healthy!")
	})

	api := app.Group("/api/v1")

	jots := api.Group("/jots")
	jotRepo := storage.NewJotRepository(db)
	jotController := ct.NewJotController(jotRepo)

	// add routes here
	jots.Post("/", jotController.Create)
	jots.Get("/", jotController.GetAll)
}
