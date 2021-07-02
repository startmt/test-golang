package driver

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	. "github.com/startmt/test-golang/test/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoResource struct {
	DB *mongo.Database
}

func ConnectMongo(ctx context.Context, conf Config) (MongoResource, error) {
	mongoConnectString := fmt.Sprintf("mongodb://%s:%s@%s", conf.Mongo.Username, conf.Mongo.Password, conf.Mongo.Host)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConnectString))
	if err != nil {
		return MongoResource{}, errors.Wrap(err, "connecting mongodb")
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return MongoResource{}, errors.Wrap(err, "pinging to mongodb")
	}

	return MongoResource{DB: client.Database(conf.Mongo.DBName)}, nil
}
