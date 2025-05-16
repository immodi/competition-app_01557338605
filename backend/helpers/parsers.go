package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type EventAssignRequest struct {
	UserID int64 `json:"userId"`
}

func ParseTheUserIdFromRequest(r *http.Request) (int64, error) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return 0, err
	}

	bodyReader1 := bytes.NewReader(bodyBytes)

	r.Body = io.NopCloser(bytes.NewReader(bodyBytes))

	var req EventAssignRequest
	if err := json.NewDecoder(bodyReader1).Decode(&req); err != nil {
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
