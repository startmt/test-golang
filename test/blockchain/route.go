package blockchain

import (
	"fmt"
	"net/http"

	"github.com/startmt/test-golang/test/constant"
)

type Router struct{}

// Get Point: -1
func (r *Router) Get(path string, handler http.HandlerFunc) {
	http.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method == constant.GET {
			handler(w, req)
		} else {
			http.Error(w, "method not allow", http.StatusMethodNotAllowed)
		}
	})
}

func (r *Router) Post(path string, handler http.HandlerFunc) {
	http.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method == constant.POST {
			handler(w, req)
			return
		}
		http.Error(w, "method not allow", http.StatusMethodNotAllowed)
		return
	})
}

func (r *Router) Handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")

	fmt.Println(req.Method, req.URL.Path)

	Server := http.DefaultServeMux
	Server.ServeHTTP(w, req)
}
