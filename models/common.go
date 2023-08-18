package models

import (
	"battery-detector/configs"

	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var chargeCollection *mongo.Collection = configs.GetCollection(configs.DB, "charges")
