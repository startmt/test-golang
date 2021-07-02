package blockchain

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

func GetBlockChainArrayController(resource DefaultControllerResource) ([]BlockChain, error) {
	blocks, err := QueryBlockChain(resource)

	if err != nil {
		return []BlockChain{}, err
	}
	return blocks, nil
}

func GetBlockChainByHashController(resource DefaultControllerResource, hash string) (BlockChain, error) {
	searchChain, err := GetBlockChainOneBy(resource)(bson.M{"hash": hash})

	if err != nil {
		return BlockChain{}, err
	}
	return searchChain, nil
}

func GetBlockChainByIndexController(resource DefaultControllerResource, index int) (BlockChain, error) {
	searchChain, err := GetBlockChainOneBy(resource)(bson.M{"index": index})

	if err != nil {
		return BlockChain{}, err
	}
	return searchChain, nil
}

func AddBlockChainController(resource DefaultControllerResource, body []byte) error {
	var reqBody CreateBlockChainRequest
	if err := json.Unmarshal(body, &reqBody); err != nil {
		return err
	}

	blocks, err := QueryBlockChain(resource)
	if err != nil {
		return err
	}

	block := BlockChain{
		Index: len(blocks),
		Body:  reqBody.Body,
	}

	serviceParam := AddPrevHashInBlock(blocks, block)
	newBlock := NewBlockByBody(serviceParam)

	if err = InsertBlockChainOne(resource, newBlock); err != nil {
		return err
	}

	return nil
}

func ValidateBlockChainController(resource DefaultControllerResource) (ValidateBlockChainResponse, error) {
	blocks, err := QueryBlockChain(resource)
	if err != nil {
		return ValidateBlockChainResponse{IsValidate: false}, err
	}

	isBlockValidate := ValidateBlockChain(blocks)
	return ValidateBlockChainResponse{IsValidate: isBlockValidate}, nil
}
