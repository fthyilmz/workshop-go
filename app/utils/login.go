package utils

import (
	"context"
	"github.com/fthyilmz/workshop-go.git/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type LoginService struct {
	Client *mongo.Client
}

func (service *LoginService) LoginUser(username string, password string) bool {

	var user model.User

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err := service.Client.Database("inventory").Collection("users").FindOne(ctx, bson.M{"username": username}).Decode(&user)

	if err != nil {
		return false
	}

	isVerifiedPassword := user.CheckPassword(password)

	if !isVerifiedPassword {
		return false
	}

	return true
}
