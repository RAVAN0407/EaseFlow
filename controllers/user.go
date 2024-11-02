package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/RAVAN0407/EaseFlow/common"
	"github.com/RAVAN0407/EaseFlow/database"
	"github.com/RAVAN0407/EaseFlow/models"
	"github.com/gin-gonic/gin"
)

func SingUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if common.ValidateUser(user) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid detials"})
			return
		}
		tc := database.OpenConnection(database.Client, models.USER)
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		result, err := tc.InsertOne(context.TODO(), user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{"successfull": result})
	}

}
