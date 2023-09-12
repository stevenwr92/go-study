package routes

import (
	"demo/handler"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	api := app.Group("/api")
	user := api.Group("/user")

	user.Get("/", handler.GetUsers)

	user.Get("/:id", handler.GetUserById)

	user.Post("/", handler.CreateUser)

	user.Put("/:id", handler.EditUser)

	user.Delete("/:id", handler.DeleteUser)

}
