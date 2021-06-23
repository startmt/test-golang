package blockchain

type BlockChain struct {
	Index    int    `json:"index"`
	Body     string `json:"body"`
	PrevHash string `json:"prevHash"`
	Hash     string `json:"hash"`
}

func (p *BlockChain) MakeBlockWithNewHash(block BlockChain) BlockChain{
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
type ChainArray []BlockChain
var chain = ChainArray{}

func (chain ChainArray) SearchBlockChainBy(fn func(blockChain BlockChain) bool) (BlockChain, error) {
	for _,block := range chain {
		if fn(block) {
			return block, nil
		}
	}
	return BlockChain{}, ErrorNotFound
}


func IsSameHash(hash string) func(blockChain BlockChain) bool {
	return func(blockChain BlockChain) bool{
		return blockChain.Hash == hash
	}
}

func IsSameIndex(index int) func(blockChain BlockChain) bool {
	return func(blockChain BlockChain) bool{
		return blockChain.Index == index
	}
}