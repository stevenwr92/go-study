package handler

import (
	"demo/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	user := []models.User{}

	models.DB.Find(&user)

	return c.Status(fiber.StatusOK).JSON(user)
}

func GetUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	user := models.User{}

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Params bukan integer",
		})
	}

	models.DB.First(&user, id)

	if user.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "User dengan id " + strconv.Itoa(id) + " tidak terdaftar.",
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	user := models.User{}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := models.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func EditUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	user := models.User{}

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Params bukan integer",
		})
	}

	models.DB.First(&user, id)

	if user.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "User dengan id " + strconv.Itoa(id) + " tidak terdaftar.",
		})
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if models.DB.Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data tidak di temukan",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User berhasil di update",
		"user":    user,
	})

}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	user := models.User{}

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Params bukan integer",
		})
	}

	if models.DB.Delete(&user, id).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User tidak di temukan",
		})
	}

	return c.JSON(fiber.Map{
		"message": " User dengan id: " + strconv.Itoa(id) + " berhasil di hapus.",
	})
}
