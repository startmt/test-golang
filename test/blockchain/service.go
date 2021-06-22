package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
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

func SearchBlockChain(array []BlockChain, hash string) BlockChain {
	for i := 0; i < len(array); i++ {
		if array[i].Hash == hash {
			return array[i]
		}
	}
	return BlockChain{}

}
