package controllers

import (
	"github.com/akhil-is-watching/techletics_alumni_reg/repository"
	"github.com/akhil-is-watching/techletics_alumni_reg/storage"
	"github.com/akhil-is-watching/techletics_alumni_reg/types"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
)

func LoginVolunteer(c *fiber.Ctx) error {
	var input types.LoginRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}

	volunteerRepository := repository.NewVolunteerRepository(storage.GetDB())
	result, err := volunteerRepository.GetByAuth(input)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"data":  "Invalid Username or Password",
		})
	}

	claims := jtoken.MapClaims{
		"id":    result.ID,
		"email": result.Email,
		"name":  result.Name,
	}

	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"data":  types.LoginResponse{Token: t, Volunteer: result},
	})
}
