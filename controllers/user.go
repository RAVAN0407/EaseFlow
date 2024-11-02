package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/RAVAN0407/EaseFlow/database"
	"github.com/RAVAN0407/EaseFlow/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

var validate = validator.New()

func SingUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = validate.Struct(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		tc := database.OpenConnection(database.Client, models.USER)
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		user.UserID = uuid.New()
		result, err := tc.InsertOne(context.TODO(), user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{"successfull": result})
	}

}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	}

}
