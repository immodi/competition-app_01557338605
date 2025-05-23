package routes

import (
	"encoding/json"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
