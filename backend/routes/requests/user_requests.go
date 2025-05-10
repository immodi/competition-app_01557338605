package requests

type UserCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRoleUpdateRequest struct {
	UserId int64  `json:"userId"`
	Role   string `json:"role"`
}
