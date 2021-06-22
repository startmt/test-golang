package blockchain

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetBlockChainArrayController(h http.ResponseWriter, req http.Request) {
	json.NewEncoder(h).Encode(chain)
}

func GetBlockChainByHashController(h http.ResponseWriter, req http.Request) {
	strPath := strings.Split(req.URL.Path, "/")
	data := chain.Search(strPath[2])
	json.NewEncoder(h).Encode(data)

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
		Index: len(chain) - 1,
		Body:  bodyStruct.Body,
	}
	if len(chain) > 1 {
		serviceParam.PrevHash = chain[len(chain)-1].PrevHash
	}
	newBlock := AddOneChain(serviceParam)

	chain.Add(newBlock)
}
