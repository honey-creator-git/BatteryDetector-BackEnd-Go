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
			Role:      "User",
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

func SignInUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var input BatteryDetectorRequest.LoginRequest
		defer cancel()

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "JSON Binding Error."})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error occurred in validating request body."})
			return
		}

		user, err := models.LoginCheck(input.Email, input.Password, c)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Checking login failed."})
			return
		}

		config, _ := configs.LoadConfig(".")
		access_token, err := utilities.CreateToken(config.AccessTokenExpiresIn, user.Email, config.AccessTokenPrivateKey)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred in generating user token.", "status": "failed"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": gin.H{"token": access_token, "data": user}})
	}
}

func GoogleAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var input BatteryDetectorRequest.GoogleAuthRequest
		defer cancel()

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "JSON Binding Error."})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error occurred in validating request body."})
			return
		}

		user, err := models.GoogleAuthCheck(input.Email, c)
		if err != nil {
			newUser := models.User{
				ID:        primitive.NewObjectID(),
				FirstName: input.FirstName,
				LastName:  input.LastName,
				Email:     input.Email,
			}
			nUser, err := newUser.SaveUser(c)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error Occurred for add a new user."})
				return
			}

			config, _ := configs.LoadConfig(".")
			access_token, err := utilities.CreateToken(config.AccessTokenExpiresIn, input.Email, config.AccessTokenPrivateKey)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error Occurred in getting user token."})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": gin.H{"data": nUser, "token": access_token}})
		} else {
			config, _ := configs.LoadConfig(".")
			access_token, err := utilities.CreateToken(config.AccessTokenExpiresIn, user.Email, config.AccessTokenPrivateKey)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured for getting user token.", "status": "failed"})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": gin.H{"token": access_token, "data": user}})
		}
	}
}

func SetLogoutWithID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var input BatteryDetectorRequest.SetLogoutWithID
		defer cancel()

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "JSON Binding Error Occurred"})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error Occurred in validating request body."})
			return
		}

		selectedUser, err := models.FindUserWithID(input.UserID, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Failed"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": selectedUser})
	}
}
