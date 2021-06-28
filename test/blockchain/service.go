package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"reflect"
)

func CreateNewHash(b BlockChain) string {
	buildBlockByte, _ := json.Marshal(b)
	sum := sha256.Sum256(buildBlockByte)
	newHash := sum[:]
	return fmt.Sprintf("%x", newHash)

}

func AddPrevHashInBlock(blocks []BlockChain, block BlockChain) BlockChain {
	if len(blocks) > 0 {
		block.PrevHash = blocks[len(blocks)-1].Hash
	}
	return block
}

func NewBlockByBody(block BlockChain) BlockChain {

	newBlock := BlockChain{
		Index:    block.Index,
		Body:     block.Body,
		PrevHash: block.PrevHash,
	}
	return BlockChain{
		Index:    newBlock.Index,
		Body:     newBlock.Body,
		Hash:     CreateNewHash(newBlock),
		PrevHash: newBlock.PrevHash,
	}
}

func AppendBlockInDatabase(collection Collection, block BlockChain) error {
	return InsertBlockChainOne(collection, block)
}

func ValidateBlockChain(blockChain []BlockChain) bool {
	for _, b := range blockChain {
		block := BlockChain{
			Index:    b.Index,
			Body:     b.Body,
			PrevHash: b.PrevHash,
		}
		newHash := CreateNewHash(block)
		if !reflect.DeepEqual(newHash, b.Hash) {
			return false
		}
	}
	return true
}

func GetAllBlockChain(collection Collection) ([]BlockChain, error) {
	return QueryBlockChain(collection)
}

func GetBlockChainBy(fn func(collection Collection) (BlockChain, error)) func(Collection) (BlockChain, error) {
	return func(collection Collection) (BlockChain, error) {
		return fn(collection)
	}
}
