package controllers

import (
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/salty-max/grest/internal/models"
	"github.com/salty-max/grest/internal/storage"
)

type JotController struct {
	Repo *storage.JotRepository
}

type CreateJotDTO struct {
	Slug   string `json:"slug"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

func NewJotController(repo *storage.JotRepository) *JotController {
	return &JotController{
		Repo: repo,
	}
}

func (j *JotController) GetAll(c *fiber.Ctx) error {
	jots, err := j.Repo.GetJots(c.Context())
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
	newJot, err := j.Repo.CreateJot(c.Context(), *jot)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newJot)
}
