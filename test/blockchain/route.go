package blockchain

import (
	"fmt"
	"github.com/startmt/test-golang/test/constant"
	"net/http"

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
			return
		}

		http.Error(w, "methodnotallow", http.StatusMethodNotAllowed)
		return
	})
}

func (r *Router) Handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Println(req.Method, req.URL.Path)
	Server := http.DefaultServeMux
	Server.ServeHTTP(w, req)
}
