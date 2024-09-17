package handlers

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Abyss-fallers/ITB-go-back/database"
	"github.com/Abyss-fallers/ITB-go-back/models"
	"gopkg.in/gomail.v2"
)

func Send(subject, body string, to []string) error {
	host := os.Getenv("MAILTRAP_HOST")
	portStr := os.Getenv("MAILTRAP_PORT")
	user := os.Getenv("MAILTRAP_USER")
	pass := os.Getenv("MAILTRAP_PASS")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return fmt.Errorf("invalid port: %v", err)
	}

	sender := os.Getenv("SENDER")

	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(host, port, user, pass)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("could not send email: %v", err)
	}

	return nil
}

func SendMail(user models.User) error {
	if result := database.DB.Db.Limit(1).Find(&user, user.ID); result.Error != nil {
		return errors.New("user does not exist")
	}

	to := []string{user.Email}
	subject := "OTP for account verification"
	body := fmt.Sprintf("Your OTP is: %s", "idi nahui")

	if err := Send(subject, body, to); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	log.Println("Email sent to:", user.Email)

	return nil
}
