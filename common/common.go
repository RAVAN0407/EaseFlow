package common

import (
	"net/mail"
	"os"

	"github.com/RAVAN0407/EaseFlow/models"
)

func GetAddress() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":" + "8080"
	}
	return ":" + port
}

func ValidateUser(user models.User) bool {
	if user.UserName == "" {
		return true
	}
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return true
	}
	if user.Password == "" {
		return true
	}
	return false
}
