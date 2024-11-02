package routers

import (
	"github.com/RAVAN0407/EaseFlow/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.GET("/login", controllers.Login())
	r.GET("/signup", controllers.SingUp())
}
