package requests

type EventRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Date        string  `json:"date"`
	Venue       string  `json:"venue"`
	Price       float64 `json:"price"`
	Image       []byte  `json:"image,omitempty"`
}

type EventAssignRequest struct {
	UserID int64 `json:"userId"`
}
