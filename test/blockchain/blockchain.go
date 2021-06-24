package blockchain

import "errors"

var (
	ErrorNotFound = errors.New("not found")
	chain         []BlockChain
)
