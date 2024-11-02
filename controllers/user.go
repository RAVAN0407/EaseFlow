package controllers

import (
	"context"
	"net/http"

	"github.com/RAVAN0407/EaseFlow/common"
	"github.com/RAVAN0407/EaseFlow/database"
	"github.com/RAVAN0407/EaseFlow/models"
	"github.com/gin-gonic/gin"
)

func SingUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		err := c.BindJSON(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		if common.ValidateUser(user) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid detials"})
		}
		tc := database.OpenConnection(database.Client, models.USER)
		result, err := tc.InsertOne(context.TODO(), user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusAccepted, gin.H{"successfull": result})
	}

}
