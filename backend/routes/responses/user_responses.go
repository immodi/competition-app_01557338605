package responses

type UserResponse struct {
	UserId    int64  `json:"userId"`
	Role      string `json:"role"`
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
	Tickets   int64  `json:"tickets"`
}

type UserDeletionResponse struct {
	Message string `json:"message"`
}
