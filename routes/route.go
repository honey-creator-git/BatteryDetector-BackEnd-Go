package routes

import (
	"battery-detector/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/api/v1/user/create", controllers.CreateUser())
	router.POST("/api/v1/user/signin", controllers.SignInUser())
}

func BatteryDetectorRoute(router *gin.Engine) {
	UserRoutes(router)
}
