package common

import (
	"os"
)

func GetAddress() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":" + "8080"
	}
	return ":" + port
}
