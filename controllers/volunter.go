package controllers

import (
	"github.com/akhil-is-watching/techletics_alumni_reg/repository"
	"github.com/akhil-is-watching/techletics_alumni_reg/storage"
	"github.com/akhil-is-watching/techletics_alumni_reg/types"
	"github.com/gofiber/fiber/v2"
)

func CreateVolunteer(c *fiber.Ctx) error {
	var input types.VolunteerCreateInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}

	volunteerRepository := repository.NewVolunteerRepository(storage.GetDB())
	if err := volunteerRepository.Create(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"data":  "Invalid Username or Password",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"data":  "Volunteer created successfully",
	})
}
