package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/RAVAN0407/EaseFlow/database"
	"github.com/RAVAN0407/EaseFlow/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task models.Tasks
		err := c.BindJSON(&task)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		taskCollection := database.OpenConnection(database.Client, models.USER)
		task.CreatedAt = time.Now()
		task.UpdatedAt = time.Now()
		task.TaskID = uuid.New()
		result, err := taskCollection.InsertOne(context.TODO(), task)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"successfull": result})

	}
}

func UpdateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task models.Tasks
		err := c.BindJSON(&task)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		taskCollection := database.OpenConnection(database.Client, models.USER)
		task.CreatedAt = time.Now()
		task.UpdatedAt = time.Now()
		task.TaskID = uuid.New()
		result, err := taskCollection.UpdateOne(context.TODO(), task)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"successfull": result})

	}
}
