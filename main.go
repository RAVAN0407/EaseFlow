package main

import (
	"github.com/RAVAN0407/EaseFlow/common"
	"github.com/RAVAN0407/EaseFlow/routers"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "hello"})
	})
	routers.AuthRoutes(r)
	r.Run(common.GetAddress())
}
