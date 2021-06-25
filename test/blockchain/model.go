package blockchain

type BlockChain struct {
	Index    int    `bson:"index" json:"index"`
	Body     string `bson:"body" json:"body"`
	PrevHash string `bson:"prev_hash" json:"prevHash"`
	Hash     string `bson:"hash" json:"hash"`
}
