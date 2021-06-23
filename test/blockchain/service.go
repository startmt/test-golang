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
	for i := 0; i < len(blockChain); i++ {
		block := BlockChain{
			Index:    blockChain[i].Index,
			Body:     blockChain[i].Body,
			PrevHash: blockChain[i].PrevHash,
		}
		newHash := CreateNewHash(block)
		if !reflect.DeepEqual(newHash, blockChain[i].Hash) {
			return false
		}
	}
	return true
}
