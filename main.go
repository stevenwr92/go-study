package main

import (
	"demo/models"
	"demo/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDatabase()

	app := fiber.New()

	routes.BookRoutes(app)
	routes.UserRoutes(app)
	fmt.Println("Server is running at port:" + "3000")
	app.Listen(":3000")
}
