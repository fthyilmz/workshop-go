package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/fthyilmz/workshop-go.git/app/model"
	"github.com/fthyilmz/workshop-go.git/config"
)

var roomRepository RoomRepository

type RoomRepository struct {
	client     *mongo.Client
	database   string
	collection string
	timeout    time.Duration
}

func init() {

	roomRepository = RoomRepository{
		config.GetMongodbClient(),
		"inventory",
		"rooms",
		10,
	}
}

func GetRoomRepository() *RoomRepository {
	return &roomRepository
}

func (r *RoomRepository) AllRoom() ([]model.Room, error) {

	ctx, _ := context.WithTimeout(context.Background(), r.timeout*time.Second)

	col, err := r.client.Database(r.database).Collection(r.collection).Find(ctx, bson.D{})
	var list []model.Room
	if err != nil {
		return nil, err
	}

	for col.Next(ctx) {
		var result bson.M
		err := col.Decode(&result)
		var item model.Room
		record, _ := json.Marshal(result)
		err = json.Unmarshal(record, &item)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}

	if err := col.Err(); err != nil {
		return nil, err
	}

	defer col.Close(ctx)

	return list, nil
}

func (r *RoomRepository) GetRoomByApartmentId(id string) ([]model.Room, error) {

	ctx, _ := context.WithTimeout(context.Background(), r.timeout*time.Second)

	var list []model.Room

	filter := bson.M{"apartment": id}

	col, err := r.client.Database(r.database).Collection(r.collection).Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for col.Next(ctx) {
		var result bson.M
		err := col.Decode(&result)
		var item model.Room
		record, _ := json.Marshal(result)
		err = json.Unmarshal(record, &item)

		if err != nil {
			return nil, err
		}

		list = append(list, item)
	}

	return list, nil
}

func (r *RoomRepository) TotalFurnitureOfApartment(rooms []model.Room) (interface{}, error) {

	ctx, _ := context.WithTimeout(context.Background(), r.timeout*time.Second)

	ids := []bson.M{}

	for _, v := range rooms {
		ids = append(ids, bson.M{"room": v.Id.Hex()})
	}

	matchStage := bson.D{{"$match", bson.M{"$or": ids}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", ""}, {"total", bson.D{{"$sum", "$price"}}}}}}

	col, err := r.client.Database(r.database).Collection("furnitures").Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})

	if err != nil {
		panic(err)
	}

	var showsWithInfo []bson.M
	if err = col.All(ctx, &showsWithInfo); err != nil {
		return nil, err
	}

	fmt.Println(showsWithInfo)

	return showsWithInfo, nil
}
