package routes

import (
	"battery-detector/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/api/v1/user/create", controllers.CreateUser())
}

func BatteryDetectorRoute(router *gin.Engine) {
	UserRoutes(router)
}
