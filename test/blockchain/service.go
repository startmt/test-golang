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

func NewBlockBy(block BlockChain) BlockChain {
		newBlock := BlockChain{
			Index: block.Index,
			Body: block.Body,
			PrevHash: block.PrevHash,
		}
		return BlockChain{
			Index: newBlock.Index,
			Body: newBlock.Body,
			Hash: CreateNewHash(newBlock),
			PrevHash: newBlock.PrevHash,
	}
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
