package main

import (
	"net/http"

	"example.com/test/blockchain"
)

func main() {
	router := blockchain.Router{}

	blockchain.MainRoute(router)

	http.ListenAndServe(":8080", http.HandlerFunc(router.CreateServer))
}
