package routes

import (
	"database/sql"
	"encoding/json"
	"immodi/submission-backend/helpers"
	"immodi/submission-backend/repos"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	UserId    int64  `json:"userId"`
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
}

func UsersRouter(r chi.Router, db *sql.DB) {
	userRepository := repos.NewUserRepository(db)

	r.Get("/", getAllUsers(userRepository))
	r.Post("/", createUser(userRepository))
	r.Get("/{id}", getUser(userRepository))

}

func getAllUsers(userRepo *repos.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := userRepo.GetAllUsers()
		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, "Failed to get all users")
			return
		}

		helpers.HttpJson(w, http.StatusOK, users)
	}
}

func createUser(userRepo *repos.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			helpers.HttpError(w, http.StatusBadRequest, "Invalid request")
			return
		}
		if req.Username == "" || req.Password == "" {
			helpers.HttpError(w, http.StatusBadRequest, "Missing username and password")
			return
		}

		userId, err := userRepo.CreateUser(req.Username, req.Password)
		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, "Could not create user")
			return
		}

		user, err := userRepo.GetUserById(userId)
		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, "User creation succeeded but fetch failed")
			return
		}

		res := &UserResponse{
			UserId:    user.ID,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
		}

		helpers.HttpJson(w, http.StatusCreated, res)
	}
}

func getUser(userRepo *repos.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			helpers.HttpError(w, http.StatusBadRequest, "Invalid user ID, pass a valid one")
			return
		}

		user, err := userRepo.GetUserById(id)
		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, "Could not retrieve user")
			return
		}
		if user == nil {
			helpers.HttpError(w, http.StatusNotFound, "User not found")
			return
		}

		response := &UserResponse{
			UserId:    user.ID,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
		}
		helpers.HttpJson(w, http.StatusOK, response)
	}
}
