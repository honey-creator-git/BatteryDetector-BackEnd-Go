package controllers

import (
	"battery-detector/BatteryDetectorRequest"
	"battery-detector/BatteryDetectorResponse"
	"battery-detector/models"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddNewCharge() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": BatteryDetectorResponse.NOT_LOGGED_IN, "status": "User is not logged in."})
			return
		}
		user_email := fmt.Sprintf("User Email %v", email)
		fmt.Printf("User Email => %v", user_email)
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var input BatteryDetectorRequest.AddCharge
		defer cancel()

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error occurred in binding JSON for request"})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error in validating request"})
			return
		}

		newCharge := models.Charge{
			ID:        primitive.NewObjectID(),
			Name:      input.Name,
			IPAddress: input.IPAddress,
			LatLon:    input.LatLon,
		}

		curCharge, err := newCharge.SaveCharge(c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred in adding a new charge"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": curCharge})
	}
}
