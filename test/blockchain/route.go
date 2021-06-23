package blockchain

import (
	"fmt"
	"net/http"

	"example.com/test/constant"
)

type Router struct{}

func (r *Router) Get(path string, handler http.HandlerFunc) {
	http.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method == constant.GET {
			handler(w, req)
		} else {
			http.Error(w, "methodnotallow", http.StatusMethodNotAllowed)
		}

	})

}

func (r *Router) Post(path string, handler http.HandlerFunc) {

	http.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {

		if req.Method == constant.POST {
			handler(w, req)
		} else {
			http.Error(w, "methodnotallow", http.StatusMethodNotAllowed)
		}

	})

}

func (r *Router) CreateServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Println(req.Method, req.URL.Path)
	Server := http.DefaultServeMux
	Server.ServeHTTP(w, req)
}

func MainRoute(router Router) {
	router.Get("/blockchain/", GetBlockChainArrayController)
	router.Post("/blockchain/add", AddBlockChainController)
	router.Get("/blockchain/hash/", GetBlockChainByHashController)
	router.Get("/blockchain/index/", GetBlockChainByIndexController)
	router.Post("/blockchain/validate", ValidateBlockChainController)
}
