package blockchain

import (
	"context"
)

type CreateBlockChainRequest struct {
	Body string `json:"body"`
}

type ValidateBlockChainResponse struct {
	IsValidate bool `json:"isValidate"`
}

type DefaultControllerResource struct {
	Context    context.Context
	Collection Collection
}
