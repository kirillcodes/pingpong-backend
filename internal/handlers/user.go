package handlers

import (
	"pingPong/internal/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error { // парсинг JSON из запроса в структуру User
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Ошибка": "Не удалось выполнить запрос JSON",
		})
	}

	if err := h.DB.Create(&user).Error; err != nil { // сохранение в бд
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Ошибка": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
