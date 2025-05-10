package responses

// type EventCreateRequest struct {
// 	Name        string    `json:"name"`
// 	Description string    `json:"description"`
// 	Category    string    `json:"category"`
// 	Date        string `json:"date"`
// 	Venue       string    `json:"venue"`
// 	Price       float64   `json:"price"`
// 	Image       []byte    `json:"image,omitempty"`
// }

type EventCreateResponse struct {
	EventId int64 `json:"eventId"`
}

type EventResponse struct {
}
