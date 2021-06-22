package blockchain

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"example.com/test/constant"
)

func GetBlockChainArrayController(h http.ResponseWriter, req http.Request) {
	json.NewEncoder(h).Encode(chain)
}

func GetBlockChainByHashController(h http.ResponseWriter, req http.Request) {
	strPath := strings.Split(strings.Trim(req.URL.Path, "/"), "/")
	data, err := chain.Search(strPath[2])
	if err != nil {
		errorResponse := constant.ErrorResponse{Status: 404, Meesage: "Not found."}
		json.NewEncoder(h).Encode(errorResponse)
		return
	}
	json.NewEncoder(h).Encode(data)

}

func GetBlockChainByIndexController(h http.ResponseWriter, req http.Request) {
	strPath := strings.Split(strings.Trim(req.URL.Path, "/"), "/")
	index, err := strconv.Atoi(strPath[2])
	if err != nil || len(chain) == 0 || index >= len(chain) {
		errorResponse := constant.ErrorResponse{Status: 404, Meesage: "Not found."}
		json.NewEncoder(h).Encode(errorResponse)
		return
	}
	json.NewEncoder(h).Encode(chain[index])

}

func AddBlockChainController(h http.ResponseWriter, req http.Request) {
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		http.Error(h, err.Error(), http.StatusBadRequest)
		return
	}
	var reqBody CreateBlockChainReq
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		http.Error(h, err.Error(), http.StatusBadRequest)
		return
	}

	bodyStruct := CreateBlockChainReq{Body: string(reqBody.Body)}
	serviceParam := BlockChain{
		Index: len(chain),
		Body:  bodyStruct.Body,
	}
	if len(chain) > 0 {
		serviceParam.PrevHash = chain[len(chain)-1].Hash
	}
	newBlock := AddOneChain(serviceParam)

	chain.Add(newBlock)
}
