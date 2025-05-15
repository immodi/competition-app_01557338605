package helpers

import (
	"encoding/json"
	"fmt"
	"immodi/submission-backend/routes/requests"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func ParseTheUserIdFromRequest(r *http.Request) (int64, error) {
	var req requests.EventAssignRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return 0, err
	}

	if req.UserID == 0 {
		return 0, fmt.Errorf("missing user id")
	}

	return req.UserID, nil
}

func ParseUserIdFromRoute(r *http.Request) (int64, error) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}
