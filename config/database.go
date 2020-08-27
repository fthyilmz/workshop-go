package config

import (
	"context"
	"fmt"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	mongodbClient     *mongo.Client
	OnceMongodbClient sync.Once
)

type MongoDB struct {
	Host    string
	Port    string
	Timeout time.Duration
}

func newMongoClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(getMongoUri())
	mongodbClient, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = mongodbClient.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return mongodbClient
}

func GetMongodbClient() *mongo.Client {
	OnceMongodbClient.Do(func() {
		mongodbClient = newMongoClient()
		err := mongodbClient.Ping(context.Background(), readpref.Primary())
		if err != nil {
			log.Fatal("Couldn't connect to the Mongodb ", err)
		} else {
			log.Info("Connected to the Mongodb ")
		}
	})
	return mongodbClient
}

func getMongoUri() string {
	fmt.Println(Global.Server)
	return fmt.Sprintf("mongodb://%s:%d", Database.Host, Database.Port)
}

func init() {
	GetMongodbClient()
}
