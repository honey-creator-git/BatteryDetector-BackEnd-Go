package models

import (
	"battery-detector/BatteryDetectorResponse"
	"battery-detector/utilities"
	"context"
	"errors"
	"fmt"
	"net/mail"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty"`
	Email     string             `json:"email,omitempty"`
	FirstName string             `json:"firstName,omitempty"`
	LastName  string             `json:"lastName,omitempty"`
	Password  string             `json:"password,omitempty"`
	Role      string             `json:"role,omitempty"`
}

func validMailAddress(address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}

func ValidateEmail(email string) (bool, error) {
	add, ok := validMailAddress(email)
	if !ok {
		return false, errors.New(BatteryDetectorResponse.EMAIL_VERIFY_ERROR)
	} else {
		fmt.Printf("Email is %v", add)
		return true, nil
	}
}

func (newUser *User) SaveUser(c context.Context) (User, error) {
	mUser, err := userCollection.InsertOne(c, newUser)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return User{}, errors.New("email already exist")
		}
	}

	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}

	if _, err := userCollection.Indexes().CreateOne(c, index); err != nil {
		return User{}, errors.New("could not create index of email")
	}

	var curUser User
	userCollection.FindOne(c, bson.M{"_id": mUser.InsertedID}).Decode(&curUser)

	return curUser, nil
}

func LoginCheck(email string, password string, c context.Context) (User, error) {
	var err error

	user := User{}

	err = userCollection.FindOne(c, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return User{}, errors.New(BatteryDetectorResponse.EMAIL_NOT_FOUND)
	}

	err = utilities.VerifiyPassword(user.Password, password)
	if err != nil {
		return user, errors.New(BatteryDetectorResponse.WRONG_PASSWORD)
	}

	return user, nil
}

func GoogleAuthCheck(email string, c context.Context) (User, error) {
	var err error

	user := User{}
	err = userCollection.FindOne(c, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return User{}, errors.New(BatteryDetectorResponse.EMAIL_NOT_FOUND)
	}

	return user, nil
}

func FindUserWithID(objId primitive.ObjectID, c context.Context) (User, error) {
	var user User
	userCollection.FindOne(c, bson.M{"id": objId}).Decode(&user)

	if user.ID.Hex() == "000000000000000000000000" {
		return User{}, errors.New(BatteryDetectorResponse.NOT_EXIST_USER)
	}

	return user, nil
}
