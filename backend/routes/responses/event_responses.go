package responses

import "immodi/submission-backend/repos"

type EventResponse struct {
	EventId int64 `json:"eventId"`
}

type EventsResponse struct {
	Events []repos.Event `json:"events"`
	Count  int           `json:"count"`
}

type EventDeletionResponse struct {
	Id      int64  `json:"id"`
	Message string `json:"message"`
}
