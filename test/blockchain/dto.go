package blockchain

import (
	"encoding/json"
	"net/http"
)

func UnMarshal(h http.ResponseWriter, data []byte, res interface{}) {
	err := json.Unmarshal(data, &res)
	if err != nil {
		http.Error(h, err.Error(), http.StatusBadRequest)
		return
	}

}
