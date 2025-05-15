package responses

type EventResponse struct {
	EventId int64 `json:"eventId"`
}

type EventDeletionResponse struct {
	Id      int64  `json:"id"`
	Message string `json:"message"`
}
