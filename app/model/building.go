package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Building struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	Title     string             `json:"title" bson:"title" validate:"required"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at" validate:"required"`
}

func NewBuilding(title string) Building {
	return Building{CreatedAt: time.Now(), Id: primitive.NewObjectID(), Title: title}
}
