package main

import (
	"net/http"

	"example.com/test/blockchain"
)

func main() {
	blockchain.MainRoute(http.HandleFunc)
	http.ListenAndServe(":8080", nil)
}
