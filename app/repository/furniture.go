package repository

import (
	"context"
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/fthyilmz/workshop-go.git/app/model"
	"github.com/fthyilmz/workshop-go.git/config"
)

var repository FurnitureRepository

type FurnitureRepository struct {
	client     *mongo.Client
	database   string
	collection string
	timeout    time.Duration
}

func init() {

	repository = FurnitureRepository{
		config.GetMongodbClient(),
		"inventory",
		"furnitures",
		10,
	}
}

func GetFurnitureRepository() *FurnitureRepository {
	return &repository
}

func (r *FurnitureRepository) All() ([]model.Furniture, error) {

	ctx, _ := context.WithTimeout(context.Background(), r.timeout*time.Second)

	col, err := r.client.Database(r.database).Collection(r.collection).Find(ctx, bson.D{})
	var list []model.Furniture
	if err != nil {
		return nil, err
	}

	for col.Next(ctx) {
		var result bson.M
		err := col.Decode(&result)
		var item model.Furniture
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

func (r *FurnitureRepository) ById(id string) (model.Furniture, error) {

	ctx, _ := context.WithTimeout(context.Background(), r.timeout*time.Second)

	var furniture model.Furniture
	objID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": objID}

	err := r.client.Database(r.database).Collection(r.collection).FindOne(ctx, filter).Decode(&furniture)

	if err != nil {
		return furniture, err
	}

	return furniture, nil
}

func (r *FurnitureRepository) Store(furniture model.Furniture) {
	ctx, _ := context.WithTimeout(context.Background(), r.timeout*time.Second)

	res, _ := r.client.Database(r.database).Collection(r.collection).InsertOne(ctx, furniture)
	if res != nil {
		log.Info("ItemId : %s", res.InsertedID)
	}

}

func (r *FurnitureRepository) Update(furniture model.Furniture) bool {

	ctx, _ := context.WithTimeout(context.Background(), r.timeout*time.Second)

	collection := r.client.Database(r.database).Collection(r.collection)

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": furniture.Id},
		bson.D{
			{"$set", bson.D{{"title", furniture.Title}, {"room", furniture.Room}, {"price", furniture.Price}}},
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	return result.ModifiedCount > 0
}

func (r *FurnitureRepository) Delete(furniture model.Furniture) bool {
	ctx, _ := context.WithTimeout(context.Background(), r.timeout*time.Second)

	collection := r.client.Database(r.database).Collection(r.collection)

	result, err := collection.DeleteOne(
		ctx,
		bson.M{"_id": furniture.Id},
	)

	if err != nil {
		log.Fatal(err)
	}

	return result.DeletedCount >= 0
}

func (r *FurnitureRepository) TotalFurniture() []bson.M {

	ctx, _ := context.WithTimeout(context.Background(), r.timeout*time.Second)

	matchStage := bson.D{{"$match", bson.D{}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", ""}, {"total", bson.D{{"$sum", "$price"}}}}}}

	col, err := r.client.Database(r.database).Collection(r.collection).Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})

	if err != nil {
		panic(err)
	}

	var showsWithInfo []bson.M
	if err = col.All(ctx, &showsWithInfo); err != nil {
		panic(err)
	}

	return showsWithInfo
}
