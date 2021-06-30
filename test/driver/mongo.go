package driver

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoResource struct {
	DB *mongo.Database
}

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
	Endpoint string `json:"endpoint"`
}

func ConnectMongo(ctx context.Context, conf Config) (MongoResource, error) {
	mongoConnectString := fmt.Sprintf("mongodb://%s:%s@%s", conf.Username, conf.Password, conf.Endpoint)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConnectString))
	if err != nil {
		return MongoResource{}, errors.Wrap(err, "connecting mongodb")
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return MongoResource{}, errors.Wrap(err, "pinging to mongodb")
	}

	return MongoResource{DB: client.Database(conf.DBName)}, nil
}
