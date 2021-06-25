package driver

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

type MongoResource struct {
	DB *mongo.Database
}

const (
	connectTimeout = 5
)

var (
	username = os.Getenv("MONGODB_USERNAME")
	password = os.Getenv("MONGODB_PASSWORD")
	dbName   = os.Getenv("MONGODB_DATABASE_NAME")
	endpoint = os.Getenv("MONGODB_ENDPOINT")
)

func ConnectMongo() (*MongoResource, error) {
	mongoConnectString := fmt.Sprintf("mongodb://%s:%s@%s", username, password, endpoint)
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConnectString))
	if err != nil {
		log.Printf("MONGODB CONNECTION ERROR:%s\n", err.Error())
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println("MONGODB PING ERROR")
		log.Println(err.Error())
		return nil, err
	}
	return &MongoResource{DB: client.Database(dbName)}, nil
}
