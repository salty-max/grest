package jot

import (
	"github.com/gofiber/fiber/v2"
)

func AddJotRoutes(app *fiber.App, controller JotController) {
	jots := app.Group("/api/v1/jots")

	// add routes here
	jots.Post("/", controller.Create)
	jots.Get("/", controller.GetAll)
}
