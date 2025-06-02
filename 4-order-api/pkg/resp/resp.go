package resp

import (
	"encoding/json"
	"log"
	"net/http"
)

func Json(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println(err)
	}
}
