package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Apartment struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	Title     string             `json:"title" bson:"title" validate:"required"`
	Building  string             `json:"building" bson:"building" validate:"required"  `
	CreatedAt time.Time          `json:"created_at" bson:"created_at" validate:"required"`
}

func NewApartment(title string, building string) Apartment {
	return Apartment{CreatedAt: time.Now(), Id: primitive.NewObjectID(), Title: title, Building: building}
}
