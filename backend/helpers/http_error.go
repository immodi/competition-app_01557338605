package helpers

import (
	"encoding/json"
	"net/http"
)

func HttpError(w http.ResponseWriter, code int, message string) {
	response := map[string]any{
		"message": message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
