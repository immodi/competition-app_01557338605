package repos

import (
	"database/sql"
	"fmt"
	"immodi/submission-backend/helpers"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (userRepository *UserRepository) GetAllUsers() ([]User, error) {
	query := "SELECT id, username, created_at FROM users"
	rows, err := userRepository.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, rows.Err()
}

// Adds a user and returns the new user's ID
func (userRepository *UserRepository) CreateUser(username string, password string) (int64, error) {
	if user, _ := userRepository.GetUserByUsername(username); user != nil {
		return 0, fmt.Errorf("user already exists")
	}

	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		return 0, err
	}

	query := "INSERT INTO users (username, password_hash) VALUES (?, ?)"

	// Using db.Exec with parameters is safe against SQL injection, trust
	result, err := userRepository.db.Exec(query, username, hashedPassword)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (userRepository *UserRepository) GetUserByUsername(username string) (*User, error) {
	var user User

	query := "SELECT id, username, created_at FROM users WHERE username = ?"

	err := userRepository.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (userRepository *UserRepository) GetUserById(id int64) (*User, error) {
	var user User

	query := "SELECT id, username, created_at FROM users WHERE id = ?"

	err := userRepository.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
