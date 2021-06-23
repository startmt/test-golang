package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

func CreateNewHash(b BlockChain) string {
	buildBlockByte, _ := json.Marshal(b)
	sum := sha256.Sum256(buildBlockByte)
	newHash := sum[:]
	return fmt.Sprintf("%x", newHash)

}

func AddOneChain(block BlockChain) BlockChain {
	block.MakeBlockWithNewHash(block)
	return block
}

func SearchBlockChain(array []BlockChain, hash string) (BlockChain, error) {
	for i := 0; i < len(array); i++ {
		if array[i].Hash == hash {
			return array[i], nil
		}
	}
	return BlockChain{}, errors.New("notfound")

}

func SearchBlockChainByPrevHash(array []BlockChain, prevHash string) (BlockChain, error) {
	for i := 0; i < len(array); i++ {
		if array[i].PrevHash == prevHash {
			return array[i], nil
		}
	}
	return BlockChain{}, errors.New("notfound")

}

func ValidateBlockChain(blockChain []BlockChain) bool {
	if len(blockChain) < 2 {
		return true
	}
	currentBlock := blockChain[0]
	for i := 1; i < len(blockChain)-1; i++ {
		block := BlockChain{
			Index:    blockChain[i].Index,
			Body:     blockChain[i].Body,
			PrevHash: blockChain[i].PrevHash,
		}
		newHash := CreateNewHash(block)
		block.Hash = newHash

		if !reflect.DeepEqual(block.Hash, blockChain[i].Hash) {
			return false
		}
		_, err := SearchBlockChainByPrevHash(blockChain[i+1:], newHash)
		if err != nil {
			return false
		}
		currentBlock = block
	}

	return reflect.DeepEqual(blockChain[len(blockChain)-1].PrevHash, currentBlock.Hash)
}
