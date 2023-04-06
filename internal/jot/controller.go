package jot

import (
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/salty-max/grest/internal/models"
)

type JotController struct {
	Store *JotStorage
}

type CreateJotDTO struct {
	Slug   string `json:"slug"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

func NewJotController(store *JotStorage) *JotController {
	return &JotController{
		Store: store,
	}
}

func (j *JotController) GetAll(c *fiber.Ctx) error {
	jots, err := j.Store.GetJots(c.Context())
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch jots",
		})
	}

	return c.JSON(jots)
}

func (j *JotController) Create(c *fiber.Ctx) error {
	// parse the request body
	jot := new(models.Jot)
	if err := c.BodyParser(&jot); err != nil {
		log.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// create the jot
	newJot, err := j.Store.CreateJot(c.Context(), *jot)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newJot)
}
