package routes

import (
	"database/sql"
	"encoding/json"
	"immodi/submission-backend/helpers"
	"immodi/submission-backend/repos"
	"immodi/submission-backend/routes/requests"
	"immodi/submission-backend/routes/responses"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func UsersRouter(r chi.Router, db *sql.DB) {
	userRepository := repos.NewUserRepository(db)

	r.Get("/", getAllUsers(userRepository))
	r.Post("/", createUser(userRepository))
	r.Put("/", updateUserRole(userRepository))
	r.Get("/{id}", getUser(userRepository))
	r.Delete("/{id}", deleteUser(userRepository))
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
		var req requests.UserCreateRequest
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

		res := &responses.UserResponse{
			UserId:    user.ID,
			Role:      user.Role,
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

		response := &responses.UserResponse{
			UserId:    user.ID,
			Role:      user.Role,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
		}
		helpers.HttpJson(w, http.StatusOK, response)
	}
}

func deleteUser(userRepo *repos.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			helpers.HttpError(w, http.StatusBadRequest, "Invalid user ID, pass a valid one")
			return
		}

		err = userRepo.DeleteUser(id)
		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, "Could not delete the user")
			return
		}

		res := &responses.UserDeletionResponse{
			Message: "User deleted successfully",
		}

		helpers.HttpJson(w, http.StatusOK, res)
	}
}

func updateUserRole(userRepo *repos.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req requests.UserRoleUpdateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			helpers.HttpError(w, http.StatusBadRequest, "Invalid request")
			return
		}
		if req.UserId == 0 || req.Role == "" {
			helpers.HttpError(w, http.StatusBadRequest, "Missing userId and role")
			return
		}

		if req.Role != "admin" && req.Role != "user" {
			helpers.HttpError(w, http.StatusBadRequest, "invalid user role, role must be 'admin' or 'user'")
			return
		}

		err := userRepo.UpdateUserRole(req.UserId, req.Role)
		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, "failed to update user role")
			return
		}

		user, err := userRepo.GetUserById(req.UserId)
		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, "User update succeeded but fetch failed")
			return
		}

		res := &responses.UserResponse{
			UserId:    user.ID,
			Role:      user.Role,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
		}

		helpers.HttpJson(w, http.StatusOK, res)
	}
}
