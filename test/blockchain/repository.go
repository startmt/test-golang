package blockchain

import (
	"context"
	"github.com/startmt/test-golang/test/cmd/driver"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Collection struct {
	Blockchain *mongo.Collection
}

const (
	connectTimeout = 20
)

func createContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	return ctx, cancel
}

func CreateCollection(resource driver.MongoResource) Collection {
	return Collection{Blockchain: resource.DB.Collection("blockchains")}
}

func QueryBlockChain(collection Collection) ([]BlockChain, error) {
	ctx, cancel := createContext()
	var blocks []BlockChain
	defer cancel()
	cur, err := collection.Blockchain.Find(ctx, bson.M{})
	if err != nil {
		return []BlockChain{}, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var block BlockChain
		var bsonBlock bson.M
		if err = cur.Decode(&bsonBlock); err != nil {
			return []BlockChain{}, err
		}
		bsonBytes, err := bson.Marshal(bsonBlock)
		if err != nil {
			return []BlockChain{}, err
		}
		err = bson.Unmarshal(bsonBytes, &block)
		if err != nil {
			return []BlockChain{}, err
		}
		blocks = append(blocks, block)
	}
	return blocks, nil
}

func QueryOneBlockChainByHash(hash string) func(collection Collection) (BlockChain, error) {
	return func(collection Collection) (BlockChain, error) {
		ctx, cancel := createContext()
		defer cancel()
		cur := collection.Blockchain.FindOne(ctx, bson.M{"hash": hash})
		var block BlockChain
		err := cur.Decode(&block)
		if err != nil {
			return BlockChain{}, err
		}
		return block, nil
	}
}

func QueryOneBlockChainByIndex(index int) func(collection Collection) (BlockChain, error) {
	return func(collection Collection) (BlockChain, error) {
		ctx, cancel := createContext()
		defer cancel()
		cur := collection.Blockchain.FindOne(ctx, bson.M{"index": index})
		var block BlockChain
		err := cur.Decode(&block)
		if err != nil {
			return BlockChain{}, err
		}
		return block, nil
	}
}

func InsertBlockChainOne(collection Collection, blockChain BlockChain) error {
	ctx, cancel := createContext()
	defer cancel()
	_, err := collection.Blockchain.InsertOne(ctx, blockChain)
	return err
}
