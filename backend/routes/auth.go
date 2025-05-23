package routes

import (
	"database/sql"
	"encoding/json"
	"immodi/submission-backend/helpers"
	"immodi/submission-backend/repos"
	"immodi/submission-backend/routes/requests"
	"immodi/submission-backend/routes/responses"
	helper_structs "immodi/submission-backend/structs"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func AuthRouter(r chi.Router, db *sql.DB, api *helper_structs.API) {
	r.Post("/login", Login(api.AuthRepo))
	r.Post("/register", Register(api.UserRepo))
}

func Login(authRepo repos.AuthInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req requests.AuthRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			helpers.HttpError(w, http.StatusBadRequest, "invalid request, likey an invalid schema")
			return
		}
		if req.Username == "" || req.Password == "" {
			helpers.HttpError(w, http.StatusBadRequest, "missing username and password")
			return
		}

		user, err := authRepo.GetAuthUserByUsername(req.Username)
		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, "failed to get user")
			return
		}
		if user == nil {
			helpers.HttpError(w, http.StatusUnauthorized, "user not found")
			return
		}

		isValidPassword := helpers.CheckPasswordHash(req.Password, user.PasswordHash)

		if !isValidPassword {
			helpers.HttpError(w, http.StatusUnauthorized, "invalid password")
			return
		}

		token, err := helpers.CreateToken(user.Username)
		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, "failed to generate token")
			return
		}

		res := &responses.AuthResponse{
			Token: token,
		}

		helpers.HttpJson(w, http.StatusCreated, res)
	}
}

func Register(userRepo repos.UserInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req requests.AuthRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			helpers.HttpError(w, http.StatusBadRequest, "invalid request, likey an invalid schema")
			return
		}
		if req.Username == "" || req.Password == "" {
			helpers.HttpError(w, http.StatusBadRequest, "missing username and password")
			return
		}

		httpStatus, err := userRepo.CreateUser(req.Username, req.Password)
		if err != nil {
			helpers.HttpError(w, int(httpStatus), err.Error())
			return
		}

		token, err := helpers.CreateToken(req.Username)
		if err != nil {
			helpers.HttpError(w, http.StatusInternalServerError, "failed to generate token")
			return
		}

		res := &responses.AuthResponse{
			Token: token,
		}

		helpers.HttpJson(w, http.StatusCreated, res)
	}
}
