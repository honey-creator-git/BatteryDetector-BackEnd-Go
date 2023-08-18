package models

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ChargeUser struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	TouchedAt string `json:"touchedAt,omitempty"`
}

type Charge struct {
	ID        primitive.ObjectID `json:"id,omitempty"`
	Name      string             `json:"name,omitempty"`
	IPAddress string             `json:"ipAddress,omitempty"`
	LatLon    string             `json:"latlon,omitempty"`
	Users     []ChargeUser       `json:"users"`
}

func (newCharge *Charge) SaveCharge(c context.Context) (Charge, error) {
	mCharge, err := chargeCollection.InsertOne(c, newCharge)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return Charge{}, errors.New("IP Address ALready Exist")
		}
	}

	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"ipAddress": 1}, Options: opt}

	if _, err := chargeCollection.Indexes().CreateOne(c, index); err != nil {
		return Charge{}, errors.New("Could not create index of ipAddress")
	}

	var curCharge Charge
	chargeCollection.FindOne(c, bson.M{"_id": mCharge.InsertedID}).Decode(&curCharge)

	return curCharge, nil
}
