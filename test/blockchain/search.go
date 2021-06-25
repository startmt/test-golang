package blockchain

func SearchBlockChainBy(fn func(BlockChain) bool) func(blocks []BlockChain) (BlockChain, error) {
	return func(blocks []BlockChain) (BlockChain, error) {
		for _, block := range blocks {
			if fn(block) {
				return block, nil
			}
		}
		return BlockChain{}, ErrorNotFound
	}
}

func IsSameHash(hash string) func(blockChain BlockChain) bool {
	return func(blockChain BlockChain) bool {
		return blockChain.Hash == hash
	}
}

func IsSameIndex(index int) func(blockChain BlockChain) bool {
	return func(blockChain BlockChain) bool {
		return blockChain.Index == index
	}
}
