package blockchain

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"example.com/test/constant"
)

func GetBlockChainArrayController(h http.ResponseWriter, req *http.Request) {
	json.NewEncoder(h).Encode(chain)
}

func GetBlockChainByHashController(h http.ResponseWriter, req *http.Request) {
	strPath := strings.Split(strings.Trim(req.URL.Path, "/"), "/")
	
	data, err := chain.Search(strPath[2])
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			errorResponse := constant.ErrorResponse{Status: 404, Meesage: "Not found."}
			json.NewEncoder(h).Encode(errorResponse)
		
			return
		}
		errorResponse := constant.ErrorResponse{Status: 400, Meesage: err.Error()}
		json.NewEncoder(h).Encode(errorResponse)
		
		return
	}
	json.NewEncoder(h).Encode(data)
}

func GetBlockChainByIndexController(h http.ResponseWriter, req *http.Request) {
	strPath := strings.Split(strings.Trim(req.URL.Path, "/"), "/")
	index, err := strconv.Atoi(strPath[2])
	if err != nil || len(chain) == 0 || index >= len(chain) {
		errorResponse := constant.ErrorResponse{Status: 404, Meesage: "Not found."}
		json.NewEncoder(h).Encode(errorResponse)
		return
	}
	json.NewEncoder(h).Encode(chain[index])
}

func ValidateBlockChainController(h http.ResponseWriter, req *http.Request) {
	isBlockValidate := ValidateBlockChain(chain)

	response := ValidateBlockChainResponse{IsValidate: isBlockValidate}

	json.NewEncoder(h).Encode(response)
}

func AddBlockChainController(h http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(h, err.Error(), http.StatusBadRequest)
		return
	}

	var reqBody CreateBlockChainRequest
	if err := json.Unmarshal(body, &reqBody); err != nil {
		http.Error(h, err.Error(), http.StatusBadRequest)
		return
	}
	
	serviceParam := BlockChain{
		Index: len(chain),
		Body:  reqBody.Body,
	}

	if len(chain) > 0 {
		serviceParam.PrevHash = chain[len(chain)-1].Hash
	}
	
	chain.Add(NewBlockBy(serviceParam))
}
