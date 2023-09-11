package main

import (
	"demo/models"
	"demo/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDatabase()

	app := fiber.New()

	routes.RouteInit(app)
	app.Listen(":3000")
}
