package handlers

import (
	"errors"
	"log"
	"regexp"
	"strings"

	"github.com/Abyss-fallers/ITB-go-back/database"
	"github.com/Abyss-fallers/ITB-go-back/models"
	"gorm.io/gorm"
)

func validateEmail(email string) bool {
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailRegexPattern)

	return re.MatchString(email)
}

func ValidateUser(user models.User) error {
	user.Email = strings.ToLower(user.Email)
	if !validateEmail(user.Email) {
		return errors.New("неверный формат почты")
	}

	result := database.DB.Db.First(&models.User{}, "email = ?", user.Email)
	log.Println(result.Error)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Println(errors.Is(result.Error, gorm.ErrRecordNotFound))
		return errors.New("пользователь с таким адресом электронной почты уже зарегистрирован")
	}

	return nil
}
