package blockchain

type BlockChain struct {
	Index    int    `json:"index"`
	Body     string `json:"body"`
	PrevHash string `json:"prevHash"`
	Hash     string `json:"hash"`
}

func (p *BlockChain) MakeBlockWithNewHash(block BlockChain) {
	p.Body = block.Body
	p.Index = block.Index
	p.PrevHash = block.PrevHash
	newHash := CreateNewHash(*p)

	p.Hash = newHash

}

type ChainArray []BlockChain

var chain = ChainArray{}

func (arr *ChainArray) Search(hash string) (BlockChain, error) {
	return SearchBlockChain(*arr, hash)
}

func (arr *ChainArray) Add(b BlockChain) {
	newChain := append(*arr, b)
	*arr = newChain
}

type CreateBlockChainReq struct {
	Body string `json:"body"`
}

func (req *CreateBlockChainReq) Create(body string) {
	req.Body = body
}
