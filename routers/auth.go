package routers

import (
	"github.com/RAVAN0407/EaseFlow/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.GET("/login", controllers.Login())
	r.GET("/signup", controllers.SingUp())
}

func TaskRoutes(r *gin.Engine) {
	r.GET("/task/create", controllers.CreateTask())
	r.GET("/task/update", controllers.UpdateTask())
	r.GET("/task/delete/:taskID", controllers.DeleteTask())
	r.GET("/task/view-all/:UserID", controllers.GetAllTasks())
}
