package BatteryDetectorRequest

import (
	"battery-detector/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignUpRequest struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GoogleAuthRequest struct {
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

type SetLogoutWithID struct {
	UserID primitive.ObjectID `json:"userId" binding:"required"`
}

type AddCharge struct {
	Name      string `json:"name" binding:"required"`
	IPAddress string `json:"ipAddress" binding:"required"`
	LatLon    string `json:"latlon" binding:"required"`
}

type UpdateCharge struct {
	Name      string              `json:"name"`
	IPAddress string              `json:"ipAddress"`
	LatLon    string              `json:"latlon"`
	Users     []models.ChargeUser `json:"users"`
}
