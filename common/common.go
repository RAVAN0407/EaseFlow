package common

import (
	"net/mail"
	"os"

	"github.com/influxdata/influxdb1-client/models"
)

func GetAddress() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":" + "8080"
	}
	return ":" + port
}

func ValidateUser(user models.User) bool {
	if user.Username == "" {
		return true
	}
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return true
	}
	if user.Phone == "" {
		return true
	}
	return false
}
