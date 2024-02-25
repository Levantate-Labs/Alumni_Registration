package routes

import (
	"github.com/akhil-is-watching/techletics_alumni_reg/controllers"
	"github.com/gofiber/fiber/v2"
)

func VolunteerRoutes(app *fiber.App) {
	app.Post("/volunteer/login", controllers.LoginVolunteer)
	app.Post("/volunteer", controllers.CreateVolunteer)
}
