package routes

import (
	"github.com/akhil-is-watching/techletics_alumni_reg/controllers"
	"github.com/gofiber/fiber/v2"
)

func AlumniRoutes(app *fiber.App) {
	app.Get("/alumni/:id", controllers.GetAlumni)
	app.Post("/alumni", controllers.CreateAlumni)
}
