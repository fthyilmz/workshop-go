package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Room struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	Title     string             `json:"title" bson:"title" validate:"required"`
	Apartment string             `json:"apartment" bson:"apartment" validate:"required"  `
	CreatedAt time.Time          `json:"created_at" bson:"created_at" validate:"required"`
}

func NewRoom(title string, apartment string) Room {
	return Room{CreatedAt: time.Now(), Id: primitive.NewObjectID(), Title: title, Apartment: apartment}
}
