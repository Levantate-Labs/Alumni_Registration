package main

import (
	"fmt"
	"log"

	"github.com/akhil-is-watching/techletics_alumni_reg/config"
	"github.com/akhil-is-watching/techletics_alumni_reg/helpers"
	"github.com/akhil-is-watching/techletics_alumni_reg/routes"
	"github.com/akhil-is-watching/techletics_alumni_reg/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
)

func init() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	helpers.InitUIDGen()
	helpers.InitS3Uploader()
	storage.ConnectDB(&config)
}

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(cors.New())
	routes.InitRoutes(app)
	config, err := config.LoadConfig()
	if err != nil {
		panic("ENV NOT LOADED")
	}
	app.Listen(fmt.Sprintf("0.0.0.0:%s", config.Port))
}
