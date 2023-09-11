package handler

import (
	"demo/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllBooks(c *fiber.Ctx) error {
	books := []models.Book{}
	// var books []models.Book
	models.DB.Find(&books)
	return c.Status(fiber.StatusOK).JSON(books)
	// return c.SendString("Hello, World!")
	// return nil

}

func GetBookById(ctx *fiber.Ctx) error {
	return nil
}

func CreateBook(c *fiber.Ctx) error {
	book := models.Book{}

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
}

func EditBook(ctx *fiber.Ctx) error {
	return nil

}

func DeleteBook(ctx *fiber.Ctx) error {
	return nil

}
