package repos

import (
	"database/sql"
	"fmt"
)

type AuthUser struct {
	ID           int64  `json:"id"`
	PasswordHash string `json:"passwordHash"`
	Username     string `json:"username"`
	CreatedAt    string `json:"createdAt"`
	Role         string `json:"role"`
}

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) GetAuthUserByUsername(username string) (*AuthUser, error) {
	var u AuthUser
	err := r.db.QueryRow(
		"SELECT id, username, role, created_at, password_hash FROM users WHERE username = ?",
		username,
	).Scan(&u.ID, &u.Username, &u.Role, &u.CreatedAt, &u.PasswordHash)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}

	return &u, nil
}
