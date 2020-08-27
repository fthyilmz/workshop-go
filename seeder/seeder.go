package main

import (
	"context"
	"fmt"
	"github.com/fthyilmz/workshop-go.git/app/model"
	"github.com/fthyilmz/workshop-go.git/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func main() {

	client := config.GetMongodbClient()

	var user model.User
	var room model.Room
	var apartment model.Apartment
	var building model.Building

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	errs := client.Database("inventory").Collection("users").Drop(ctx)

	if errs != nil {
		panic(errs)
	}

	err := client.Database("inventory").Collection("users").FindOne(ctx, bson.M{"username": "admin"}).Decode(&user)

	if err == nil {
		panic(err)
	}

	user = model.NewUser("admin", "123456")

	res, err := client.Database("inventory").Collection("users").InsertOne(ctx, user)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success %v", res.InsertedID)

	building = model.NewBuilding("test-building")

	build_res, err := client.Database("inventory").Collection("buildings").InsertOne(ctx, building)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success %v", build_res.InsertedID)

	apartment = model.NewApartment("test-apartment", build_res.InsertedID.(primitive.ObjectID).Hex())

	build_apart, err := client.Database("inventory").Collection("apartments").InsertOne(ctx, apartment)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success %v", build_apart.InsertedID)

	room = model.NewRoom("test-room", build_apart.InsertedID.(primitive.ObjectID).Hex())

	build_room, err := client.Database("inventory").Collection("rooms").InsertOne(ctx, room)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success %v", build_room.InsertedID)

	room = model.NewRoom("test-room - 2", build_apart.InsertedID.(primitive.ObjectID).Hex())

	build_room2, err := client.Database("inventory").Collection("rooms").InsertOne(ctx, room)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success %v", build_room2.InsertedID)
}
