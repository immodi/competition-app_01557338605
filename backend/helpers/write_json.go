package helpers

import (
	"encoding/json"
	"net/http"
)

func HttpJson(w http.ResponseWriter, status int, object any) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(object); err != nil {
		HttpError(w, http.StatusInternalServerError, "Internal server error, failed to encode the response for some reason")
		return
	}
}
