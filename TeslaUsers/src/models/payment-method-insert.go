package models

type PaymentMethodInsert struct {
	EMail         string `json:"email" bson:"email"`
	PaymentMethod string `json:"payment_method" bson:"payment_method"`
}