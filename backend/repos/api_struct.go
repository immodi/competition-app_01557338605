package repos

type API struct {
	EventRepo *EventRepository
	UserRepo  *UserRepository
	AuthRepo  *AuthRepository
}
