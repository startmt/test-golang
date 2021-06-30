package blockchain

import (
	"github.com/startmt/test-golang/test/driver"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Collection struct {
	Blockchain *mongo.Collection
}

func CreateCollection(resource driver.MongoResource) Collection {
	return Collection{Blockchain: resource.DB.Collection("blockchains")}
}

func QueryBlockChain(resource DefaultControllerResource) ([]BlockChain, error) {
	var blocks []BlockChain

	cur, err := resource.Collection.Blockchain.Find(resource.Context, bson.M{})
	if err != nil {
		return []BlockChain{}, err
	}
	defer cur.Close(resource.Context)

	for cur.Next(resource.Context) {
		var (
			block     BlockChain
			bsonBlock bson.M
		)

		if err = cur.Decode(&bsonBlock); err != nil {
			return []BlockChain{}, err
		}

		bsonBytes, err := bson.Marshal(bsonBlock)
		if err != nil {
			return []BlockChain{}, err
		}

		if err := bson.Unmarshal(bsonBytes, &block); err != nil {
			return []BlockChain{}, err
		}
		blocks = append(blocks, block)
	}
	return blocks, nil
}

func GetBlockChainOneBy(queryData bson.M) func(DefaultControllerResource) (BlockChain, error) {
	return func(resource DefaultControllerResource) (BlockChain, error) {
		cur := resource.Collection.Blockchain.FindOne(resource.Context, queryData)
		var block BlockChain
		err := cur.Decode(&block)

		if err != nil {
			return BlockChain{}, err
		}
		return block, nil
	}
}

func InsertBlockChainOne(resource DefaultControllerResource, blockChain BlockChain) error {
	_, err := resource.Collection.Blockchain.InsertOne(resource.Context, blockChain)
	return err
}
