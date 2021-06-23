package main

import (
	"github.com/startmt/test-golang/test/blockchain"
	"net/http"
)

func main() {
	router := blockchain.Router{}

	blockchain.MainRoute(router)

	http.ListenAndServe(":8080", http.HandlerFunc(router.Handler))
}
