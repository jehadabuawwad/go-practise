package main

import (
	"theSiS/web-service-fiber/configs"
	"theSiS/web-service-fiber/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.UserRoute(app)

	configs.ConnectDB()

	app.Listen(":8000")
}
