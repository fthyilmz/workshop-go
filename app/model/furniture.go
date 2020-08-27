package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type FurnitureForm struct {
	Title string  `json:"title" bson:"title" validate:"required"`
	Room  string  `json:"room" bson:"room" validate:"required"  `
	Price float32 `json:"price"  bson:"price" validate:"required" `
}

type Furniture struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	Title     string             `json:"title" bson:"title" validate:"required"`
	Room      string             `json:"room" bson:"room" validate:"required"  `
	Price     float32            `json:"price"  bson:"price" validate:"required" `
	CreatedAt time.Time          `json:"created_at" bson:"created_at" validate:"required"`
}

func NewFurniture() Furniture {
	return Furniture{CreatedAt: time.Now(), Id: primitive.NewObjectID()}
}
