package routes

import (
	"immodi/submission-backend/helpers"
	"net/http"
)

func NotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(405)
	helpers.HttpError(w, http.StatusMethodNotAllowed, "Method is not valid")
}
