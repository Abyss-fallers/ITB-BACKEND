package handlers

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/Abyss-fallers/ITB-go-back/database"
	"github.com/Abyss-fallers/ITB-go-back/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID            uint
	UserRole          models.Role
	VerificationToken string
	jwt.RegisteredClaims
}

func GenerateVerificationToken() string {
	src := rand.NewSource(time.Now().Unix())

	return strconv.Itoa(int(100000 + src.Int63()*900000))
}

func GenerateTokenAndSetCookie(user models.User) (fiber.Cookie, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	user.VerificationToken = GenerateVerificationToken()

	result := database.DB.Db.Save(&user)
	if result.Error != nil {
		return fiber.Cookie{}, result.Error
	}

	claims := &Claims{
		UserID:            user.ID,
		UserRole:          models.RoleMember,
		VerificationToken: user.VerificationToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return fiber.Cookie{}, err
	}

	cookie := fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(7 * time.Hour),
		HTTPOnly: true,
		SameSite: "strict",
	}

	return cookie, nil
}

func DecodeToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return os.Getenv("JWT_SECRET"), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
