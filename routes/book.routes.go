package routes

import (
	"demo/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	api := app.Group("/api")
	book := api.Group("/books")

	book.Get("/", handler.GetAllBooks)

	book.Get("/:id", handler.GetBookById)

	book.Post("/", handler.CreateBook)

	book.Put("/:id", handler.EditBook)

	book.Delete("/:id", handler.DeleteBook)
}
