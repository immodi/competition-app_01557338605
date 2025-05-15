package helper_structs

import "immodi/submission-backend/repos"

type API struct {
	EventRepo *repos.EventRepository
	UserRepo  *repos.UserRepository
	AuthRepo  *repos.AuthRepository
}
