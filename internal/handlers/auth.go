package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"pingPong/internal/mail"
	"pingPong/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req RegisterRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Ошибка": "Неверный ввод"})
		}

		var existing models.User
		if err := db.Where("email = ?", req.Email).First(&existing).Error; err == nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Ошибка": "Адрес электронной почты уже зарегистрирован"})
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"Ошибка": "Не удалось хешировать пароль"})
		}

		rand.Seed(time.Now().UnixNano())
		code := fmt.Sprintf("%06d", rand.Intn(1000000))

		user := models.User{
			Email:          req.Email,
			Username:       req.Username,
			PasswordHash:   string(hash),
			EmailConfirmed: false,
			ConfirmCode:    code,
		}

		if err := db.Create(&user).Error; err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Ошибка": "Ошибка сохранения профиля"})
		}

		if err := mail.SendConfirmation(req.Email, code); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"Ошибка": "Ошибка при отправления кода на почту"})
		}

		return c.JSON(fiber.Map{"Cообщение": "На почту отправлен код подтверджения"})
	}
}
