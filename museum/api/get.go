package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"maxklammer.com/go/museum/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := data.GetAll()

	// api/exhibitions?id=34
	id := r.URL.Query()["id"]

	if id != nil {
		finalId, err := strconv.Atoi(id[0])
		if err != nil || finalId > len(res) {
			http.Error(w, "Invalid Exhibition", http.StatusBadRequest)
			return
		}

		for _, item := range res {
			if item.Id == finalId {
				json.NewEncoder(w).Encode(item)
				return
			}
		}
	} else {
		json.NewEncoder(w).Encode(data.GetAll())
	}
}
