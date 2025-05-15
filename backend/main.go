package main

import (
	"immodi/submission-backend/db"
	"immodi/submission-backend/repos"
	"immodi/submission-backend/routes"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Connect to database
	db, err := db.NewDatabase("file:db/api.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create the router
	r := chi.NewRouter()
	api := &repos.API{
		EventRepo: repos.NewEventRepository(db.DB),
		UserRepo:  repos.NewUserRepository(db.DB),
		AuthRepo:  repos.NewAuthRepository(db.DB),
	}

	// Middlewares
	r.Use(middleware.Logger)

	// Routes
	r.Get("/", routes.Root)
	r.Route("/auth", func(r chi.Router) {
		routes.AuthRouter(r, db.DB, api)
	})
	r.Route("/users", func(r chi.Router) {
		routes.UsersRouter(r, db.DB, api)
	})
	r.Route("/events", func(r chi.Router) {
		routes.EventsRouter(r, db.DB, api)
	})

	r.NotFound(routes.NotFound)
	r.MethodNotAllowed(routes.NotAllowed)

	println("Listening on port http://localhost:3000/")
	http.ListenAndServe(":3000", r)
}
