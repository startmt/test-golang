package blockchain

type BlockChain struct {
	Index    int    `json:"index"`
	Body     string `json:"body"`
	PrevHash string `json:"prevHash"`
	Hash     string `json:"hash"`
}
