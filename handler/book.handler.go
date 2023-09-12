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

}

func GetBookById(c *fiber.Ctx) error {
	id := c.Params("id")
	book := models.Book{}

	models.DB.First(&book, id)
	if book.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Book with Id " + id + " Not found",
			"data":    book,
		})
	}
	return c.JSON(book)
}

func CreateBook(c *fiber.Ctx) error {
	book := models.Book{}

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := models.DB.Create(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(book)
}

func EditBook(c *fiber.Ctx) error {
	id := c.Params("id")
	book := models.Book{}

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if models.DB.Where("id = ?", id).Updates(&book).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data tidak di temukan",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Data Berhasil di update",
		"data":    book,
	})

}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	book := models.Book{}
	if models.DB.Delete(&book, id).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Book tidak di temukan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Book id " + id + " berhasil di hapus.",
	})

}
