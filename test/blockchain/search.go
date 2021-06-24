package blockchain

func SearchBlockChainBy(fn func(BlockChain) bool) func(blockChainSlice []BlockChain) (BlockChain, error) {
	return func(blockChainSlice []BlockChain) (BlockChain, error) {
		for _,block := range blockChainSlice {
			if fn(block) {
				return block, nil
			}
		}
		return BlockChain{}, ErrorNotFound
	}
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