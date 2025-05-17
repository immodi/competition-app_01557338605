package main

import (
	"log"
	"net/http"
	"os"

	"immodi/submission-backend/db"
	"immodi/submission-backend/repos"
	"immodi/submission-backend/routes"
	helper_structs "immodi/submission-backend/structs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Ensure /db directory exists
	if err := os.MkdirAll("db", os.ModePerm); err != nil {
		log.Fatalf("Failed to create db directory: %v", err)
	}

	// Connect to database
	db, err := db.NewDatabase("file:db/api.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create the router
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "ngrok-skip-browser-warning"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	api := &helper_structs.API{
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

	log.Println("Listening on port http://localhost:8020/")
	if err := http.ListenAndServe(":8020", r); err != nil {
		log.Fatalf("HTTP server failed: %v", err)
	}
}
