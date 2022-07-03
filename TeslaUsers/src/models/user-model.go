package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Id = primitive.ObjectID

type User struct {
	ID             Id       `json:"id" bson:"_id"`
	EMail          string   `json:"email" bson:"email"`
	UserName       string   `json:"username" bson:"username"`
	Name           string   `json:"name" bson:"name"`
	LastName       string   `json:"last_name" bson:"last_name"`
	BirthDay       string   `json:"birth_day" bson:"birth_day"`
	Country        string   `json:"country" bson:"country"`
	Language       string   `json:"language" bson:"language"`
	PaymentMethods []string `json:"payment_methods" bson:"payment_methods"`
}
