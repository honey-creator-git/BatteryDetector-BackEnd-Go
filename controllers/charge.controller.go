package controllers

import (
	"battery-detector/BatteryDetectorRequest"
	"battery-detector/BatteryDetectorResponse"
	"battery-detector/models"
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getAttr(obj interface{}, fieldName string) reflect.Value {
	pointToStruct := reflect.ValueOf(obj) // addressable
	curStruct := pointToStruct.Elem()
	if curStruct.Kind() != reflect.Struct {
		panic("not struct")
	}
	curField := curStruct.FieldByName(fieldName)
	if !curField.IsValid() {
		panic("not found:" + fieldName)
	}

	return curField
}

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

func UpdateCharge() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": BatteryDetectorResponse.NOT_LOGGED_IN, "status": "User is not lgged in."})
			return
		}
		user_email := fmt.Sprintf("User Email %v", email)
		fmt.Printf("User Email => %v", user_email)
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		chargeId := ctx.Param("chargeId")
		var chargeType BatteryDetectorRequest.UpdateCharge
		defer cancel()

		objChargeId, _ := primitive.ObjectIDFromHex(chargeId)

		if err := ctx.ShouldBindJSON(&chargeType); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error occurred in JSON Binding."})
			return
		}

		if validationErr := validate.Struct(&chargeType); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error occurred in validation of request."})
			return
		}

		result, err := models.GetChargeWithID(objChargeId, c)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Could not get charge with ID."})
			return
		}

		if result.ID.Hex() == "000000000000000000000000" {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "The Charge is not existed"})
			return
		}

		t := reflect.TypeOf(chargeType)
		names := make([]string, t.NumField())
		for i := range names {
			pX := getAttr(&chargeType, t.Field(i).Name)
			fmtedPX := fmt.Sprintf("%v", pX)
			if (len(fmtedPX) != 0) && (len(fmtedPX) != 2) {
				names[i] = t.Field(i).Name
				updateCharge := bson.M{
					strings.ToLower(names[i]): pX.Interface(),
				}
				result, err := models.UpdateChargeWithID(objChargeId, updateCharge, c)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred in updating charge"})
					return
				}
				fmt.Printf("Result of Update charge => %v", result)
			}
		}
		updateCharge, errr := models.GetChargeWithID(objChargeId, c)
		if errr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Could not get updated charge with id."})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": updateCharge})
	}
}
