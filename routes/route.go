package routes

import (
	"battery-detector/controllers"

	"battery-detector/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/api/v1/user/create", controllers.CreateUser())
	router.POST("/api/v1/user/signin", controllers.SignInUser())
	router.POST("/api/v1/user/social/google_auth", controllers.GoogleAuth())
	router.POST("/api/v1/set_logout", controllers.SetLogoutWithID())
}

func ChargeRoutes(router *gin.Engine) {
	router.POST("/api/v1/charge/add", controllers.AddNewCharge())
	router.POST("/api/v1/charge/edit/:chargeId", controllers.UpdateCharge())
}

func BatteryDetectorRoute(router *gin.Engine) {
	UserRoutes(router)
	router.Use(middlewares.DeserializeUser())
	ChargeRoutes(router)
}
