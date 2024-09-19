package handlers

import (
	"log"
	"time"

	"github.com/Abyss-fallers/ITB-go-back/database"
	"github.com/Abyss-fallers/ITB-go-back/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var timeLayout = "2006-01-02T03:04:05"

func Registration(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	var user models.User

	err := c.BodyParser(&user)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(map[string]any{"error in parser": err.Error()})
	}

	log.Println(user)

	if err = ValidateUser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]any{"error": err.Error()})
	}

	user.Password, err = HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]any{"error": err.Error()})
	}

	jwtToken := jwt.New(jwt.SigningMethodHS256)

	user.VerificationToken, err = jwtToken.SignedString([]byte(GenerateVerificationToken()))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]any{"error in token": err.Error()})
	}

	user.VerificationTokenExpiresAt = time.Now()
	user.LastLoginAt = time.Now()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	result := database.DB.Db.Save(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]any{"error in db": result.Error.Error()})
	}

	if err := SendMail(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]any{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]any{"success": true, "message": "Успешная регистрация", "user": user})
}

func Authentication(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	var (
		user     models.User
		userAuth models.User
	)

	err := c.BodyParser(&userAuth)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]any{"error": err.Error()})
	}

	result := database.DB.Db.First(&user, "email = ?", userAuth.Email)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(map[string]any{"error": "пользователь не найден"})
	}

	if !CheckPasswordHash(user.Password, userAuth.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]any{"error": "пароль неверный"})
	}

	token, err := GenerateTokenAndSetCookie(int(user.ID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]any{"error": err.Error()})
	}

	c.Cookie(&token)

	return c.Status(fiber.StatusOK).JSON(map[string]any{})
}

func Authorization(c *fiber.Ctx) {

}
