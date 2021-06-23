package blockchain

type CreateBlockChainRequest struct {
	Body string `json:"body"`
}

func (req *CreateBlockChainRequest) Create(body string) {
	req.Body = body
}

type ValidateBlockChainResponse struct {
	IsValidate bool `json:"isValidate"`
}
