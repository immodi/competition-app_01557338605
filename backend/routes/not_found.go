package routes

import (
	"immodi/submission-backend/helpers"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	helpers.HttpError(w, http.StatusNotFound, "Route does not exist")
}
