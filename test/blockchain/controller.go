package blockchain

import (
	"encoding/json"
	"errors"
	"github.com/startmt/test-golang/test/constant"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func GetBlockChainArrayController(h http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(h).Encode(chain)
}

func GetBlockChainByHashController(h http.ResponseWriter, req *http.Request) {
	strPath := strings.Split(strings.Trim(req.URL.Path, "/"), "/")

	searchChain, err := SearchBlockChainBy(IsSameHash(strPath[2]))(chain)
	if err != nil {
		if errors.Is(err, ErrorNotFound) {
			errorResponse := constant.ErrorResponse{Status: 404, Meesage: "Not found."}
			json.NewEncoder(h).Encode(errorResponse)
			return
		}
			errorResponse := constant.ErrorResponse{Status: 400, Meesage: err.Error()}
			json.NewEncoder(h).Encode(errorResponse)
	}
	json.NewEncoder(h).Encode(searchChain)
}

func GetBlockChainByIndexController(h http.ResponseWriter, req *http.Request) {
	strPath := strings.Split(strings.Trim(req.URL.Path, "/"), "/")
	index, err := strconv.Atoi(strPath[2])
	if err != nil || len(chain) == 0 || index >= len(chain) {
		errorResponse := constant.ErrorResponse{Status: 404, Meesage: "Not found."}
		json.NewEncoder(h).Encode(errorResponse)
		return
	}

	searchChain, err := SearchBlockChainBy(IsSameIndex(index))(chain)
	if err != nil {
		if errors.Is(err, ErrorNotFound) {
			errorResponse := constant.ErrorResponse{Status: 404, Meesage: "Not found."}
			json.NewEncoder(h).Encode(errorResponse)
			return
		}
		errorResponse := constant.ErrorResponse{Status: 400, Meesage: err.Error()}
		json.NewEncoder(h).Encode(errorResponse)
	}
	json.NewEncoder(h).Encode(searchChain)
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
	chain = append(chain,NewBlockBy(serviceParam))
}


func ValidateBlockChainController(h http.ResponseWriter, _ *http.Request) {
	isBlockValidate := ValidateBlockChain(chain)

	response := ValidateBlockChainResponse{IsValidate: isBlockValidate}

	json.NewEncoder(h).Encode(response)
}