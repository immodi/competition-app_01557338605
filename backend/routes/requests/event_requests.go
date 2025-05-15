package requests

import "immodi/submission-backend/repos"

type EventRequest struct {
	Name         string                   `json:"name"`
	Description  string                   `json:"description"`
	Category     string                   `json:"category"`
	Date         string                   `json:"date"`
	Venue        string                   `json:"venue"`
	Price        float64                  `json:"price"`
	Image        []byte                   `json:"image,omitempty"`
	Translations []repos.EventTranslation `json:"translations"`
}

type EventAssignRequest struct {
	UserID int64 `json:"userId"`
}
