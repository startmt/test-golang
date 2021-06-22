package blockchain

import (
	"net/http"
	"strings"

	"example.com/test/constant"
)

func HandlerBlockChainPath(h http.ResponseWriter, req *http.Request) {
	h.Header().Set("content-type", "application/json")
	method := req.Method

	switch method {
	case constant.GET:
		GetMethod(h, *req)
		return
	case constant.POST:
		PostMethod(h, *req)
		return
	}
}

func GetMethod(h http.ResponseWriter, req http.Request) {
	strPath := strings.Split(strings.Trim(req.URL.Path, "/"), "/")
	if len(strPath) >= 2 {
		switch strPath[1] {
		case "hash":
			GetBlockChainByHashController(h, req)
		case "index":
			GetBlockChainByIndexController(h, req)
		default:
			http.Error(h, "", http.StatusNotFound)
		}

	} else {
		GetBlockChainArrayController(h, req)
	}

}

func PostMethod(h http.ResponseWriter, req http.Request) {
	AddBlockChainController(h, req)
}

func MainRoute(handleFunc func(pattern string, handler func(h http.ResponseWriter, req *http.Request))) {
	handleFunc("/blockchain/", HandlerBlockChainPath)
	handleFunc("/blockchain/hash", HandlerBlockChainPath)
	handleFunc("/blockchain/index", HandlerBlockChainPath)
}
