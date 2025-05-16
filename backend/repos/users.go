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
	Tickets   int64  `json:"tickets"`
	Role      string `json:"role"`
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUsers() ([]User, error) {
	rows, err := r.db.Query("SELECT id, username, role, tickets, created_at FROM users")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %w", err)
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Username, &u.Role, &u.Tickets, &u.CreatedAt); err != nil {
			return nil, fmt.Errorf("error scanning user row: %w", err)
		}
		users = append(users, u)
	}

	return users, rows.Err()
}

func (r *UserRepository) CreateUser(username, password string) (int64, error) {
	// Check for existing user
	if existing, _ := r.GetUserByUsername(username); existing != nil {
		return 400, fmt.Errorf("user '%s' already exists", username)
	}

	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		return 500, fmt.Errorf("failed to hash password: %w", err)
	}

	result, err := r.db.Exec(
		"INSERT INTO users (username, password_hash) VALUES (?, ?)",
		username, hashedPassword,
	)
	if err != nil {
		return 500, fmt.Errorf("failed to create user: %w", err)
	}

	return result.LastInsertId()
}

func (r *UserRepository) GetUserByUsername(username string) (*User, error) {
	var u User
	err := r.db.QueryRow(
		"SELECT id, username, role, tickets, created_at FROM users WHERE username = ?",
		username,
	).Scan(&u.ID, &u.Username, &u.Role, &u.Tickets, &u.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}

	return &u, nil
}

func (r *UserRepository) GetUserById(id int64) (*User, error) {
	var u User
	err := r.db.QueryRow(
		"SELECT id, username, role, tickets, created_at FROM users WHERE id = ?",
		id,
	).Scan(&u.ID, &u.Username, &u.Role, &u.Tickets, &u.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	return &u, nil
}

func (r *UserRepository) DeleteUser(id int64) error {
	result, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete user id %d: %w", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("couldn't verify deletion result: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id %d", id)
	}

	return nil
}

func (r *UserRepository) UpdateUserRole(id int64, role string) error {
	_, err := r.db.Exec(
		"UPDATE users SET role = ? WHERE id = ?",
		role, id,
	)
	if err != nil {
		return fmt.Errorf("failed to update user id %d: %w", id, err)
	}
	return nil
}

func (r *UserRepository) IsAdmin(username string) bool {
	user, err := r.GetUserByUsername(username)
	if err != nil {
		return false
	}

	if user == nil {
		return false
	}
	return user.Role == "admin"
}

func (r *UserRepository) IsSameUser(username string, resourceOwnerId int64) bool {
	user, err := r.GetUserById(resourceOwnerId)

	if err != nil {
		return false
	}

	if user == nil {
		return false
	}

	return user.Username == username
}

func (r *UserRepository) RemoveOneTicketFromUser(id int64) error {
	_, err := r.db.Exec(
		"UPDATE users SET tickets = tickets - 1 WHERE id = ?",
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to update user id %d: %w", id, err)
	}

	return nil
}
