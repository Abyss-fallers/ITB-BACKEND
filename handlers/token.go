package handlers

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int
	jwt.RegisteredClaims
}

func GenerateVerificationToken() string {
	src := rand.NewSource(time.Now().Unix())

	return strconv.Itoa(int(100000 + src.Int63()*900000))
}

func GenerateTokenAndSetCookie(userID int) (fiber.Cookie, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		UserID: userID,
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
