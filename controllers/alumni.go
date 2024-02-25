package controllers

import (
	"fmt"
	"path/filepath"

	"github.com/akhil-is-watching/techletics_alumni_reg/helpers"
	"github.com/akhil-is-watching/techletics_alumni_reg/repository"
	"github.com/akhil-is-watching/techletics_alumni_reg/storage"
	"github.com/akhil-is-watching/techletics_alumni_reg/types"
	"github.com/gofiber/fiber/v2"
)

func CreateAlumni(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}

	files := form.File["file"]
	name := form.Value["name"][0]
	email := form.Value["email"][0]
	phone_no := form.Value["phone_no"][0]
	passout_year := form.Value["passout_year"][0]

	var randomName string

	for _, file := range files {

		src, err := file.Open()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		defer src.Close()

		randomName = helpers.UIDGen().GenerateID("FU") + filepath.Ext(file.Filename)

		err = helpers.GetS3Uploader().Upload(randomName, src)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
	}

	alumniRepo := repository.NewAlumniRepository(storage.GetDB())
	alumniID := helpers.UIDGen().GenerateID("AL")
	err = alumniRepo.Create(types.AlumniCreateInput{
		ID:          alumniID,
		Name:        name,
		Email:       email,
		PhoneNo:     phone_no,
		ImageURL:    fmt.Sprintf("https://techleticsassetbucket.s3.ap-south-1.amazonaws.com/%s", randomName),
		PassoutYear: passout_year,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"data":  err.Error(),
		})
	}

	go helpers.SendMail(email, fmt.Sprintf("https://alumniregistration-production.up.railway.app/alumni/%s", alumniID))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"data":  "Alumni registered successfully",
	})
}

func GetAlumni(c *fiber.Ctx) error {
	ID := c.Params("id")

	alumniRepo := repository.NewAlumniRepository(storage.GetDB())
	alumni, err := alumniRepo.Get(ID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"data":  "Registration not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"data":  alumni,
	})
}
