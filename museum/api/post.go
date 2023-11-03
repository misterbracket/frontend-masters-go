package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"maxklammer.com/go/museum/data"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		var exhibition data.Exhibition

		err := json.NewDecoder(r.Body).Decode(&exhibition)
		fmt.Println(exhibition)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data.Add(data.Exhibition{
			Id:          len(data.GetAll()) + 1,
			Title:       exhibition.Title,
			Description: exhibition.Description,
			Image:       exhibition.Image,
		})
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("OK"))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
