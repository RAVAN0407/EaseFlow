package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/RAVAN0407/EaseFlow/database"
	"github.com/RAVAN0407/EaseFlow/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
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
		filter := bson.M{"taskid": task.TaskID}
		result, err := taskCollection.UpdateOne(context.TODO(), filter, task)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"successfull": result})

	}
}

func DeleteTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		taskID := c.Param("taskId") // Assuming taskId will be passed as URL parameter
		if taskID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "task ID is required"})
			return
		}

		taskCollection := database.OpenConnection(database.Client, models.USER)
		filter := bson.M{"taskid": taskID}

		result, err := taskCollection.DeleteOne(context.TODO(), filter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if result.DeletedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
	}
}
