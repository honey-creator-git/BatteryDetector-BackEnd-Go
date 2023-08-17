package controllers

import (
	"battery-detector/BatteryDetectorRequest"
	"battery-detector/configs"
	"battery-detector/models"
	"battery-detector/utilities"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var input BatteryDetectorRequest.SignUpRequest
		defer cancel()

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error occurred in binding JSON for request"})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error Occurred in validating request"})
			return
		}

		if _, err := models.ValidateEmail(input.Email); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error occurred in validating email"})
			return
		}

		hashedPassword, _ := utilities.HashPassword(input.Password)
		newUser := models.User{
			ID:        primitive.NewObjectID(),
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Email:     input.Email,
			Password:  hashedPassword,
		}

		curUser, err := newUser.SaveUser(c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred in adding a new user"})
			return
		}

		config, _ := configs.LoadConfig(".")
		access_token, err := utilities.CreateToken(config.AccessTokenExpiresIn, input.Email, config.AccessTokenPrivateKey)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error ocurred in generating user token."})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": gin.H{"data": curUser, "token": access_token}})
		// ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": gin.H{"data": curUser}})
	}
}
