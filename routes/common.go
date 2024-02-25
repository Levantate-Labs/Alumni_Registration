package routes

import (
	"github.com/gofiber/fiber/v2"
)

func CommonRoutes(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data":  "Still Alive",
			"error": false,
		})
	})

}
