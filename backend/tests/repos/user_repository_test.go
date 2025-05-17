package tests

import (
	"database/sql"
	"immodi/submission-backend/helpers"
	"immodi/submission-backend/repos"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewUserRepository(db)

	rows := sqlmock.NewRows([]string{"id", "username", "role", "tickets", "created_at"}).
		AddRow(int64(1), "user1", "admin", int64(3), "2025-05-17T10:00:00Z").
		AddRow(int64(2), "user2", "user", int64(1), "2025-05-16T09:00:00Z")

	mock.ExpectQuery("SELECT id, username, role, tickets, created_at FROM users").
		WillReturnRows(rows)

	users, err := repo.GetAllUsers()

	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, "user1", users[0].Username)
	assert.Equal(t, int64(3), users[0].Tickets)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetUserByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewUserRepository(db)

	username := "user1"
	rows := sqlmock.NewRows([]string{"id", "username", "role", "tickets", "created_at"}).
		AddRow(int64(1), username, "admin", int64(5), "2025-05-17T10:00:00Z")

	mock.ExpectQuery("SELECT id, username, role, tickets, created_at FROM users WHERE username = ?").
		WithArgs(username).
		WillReturnRows(rows)

	user, err := repo.GetUserByUsername(username)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, username, user.Username)
	assert.Equal(t, "admin", user.Role)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestCreateUser_UserAlreadyExists(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewUserRepository(db)

	username := "user1"
	password := "pass123"

	// Mock GetUserByUsername returns a user (exists)
	rows := sqlmock.NewRows([]string{"id", "username", "role", "tickets", "created_at"}).
		AddRow(int64(1), username, "user", int64(1), "2025-05-17T10:00:00Z")
	mock.ExpectQuery("SELECT id, username, role, tickets, created_at FROM users WHERE username = ?").
		WithArgs(username).
		WillReturnRows(rows)

	id, err := repo.CreateUser(username, password)

	assert.Error(t, err)
	assert.Equal(t, int64(400), id)
	assert.Contains(t, err.Error(), "already exists")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestCreateUser_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewUserRepository(db)

	username := "newuser"
	password := "pass123"
	mockHash := "$2a$10$mockedhashedpassword"

	// Save original function to restore later
	origHashFunc := helpers.HashPassword
	defer func() { helpers.HashPassword = origHashFunc }()

	// Override HashPassword with a mock version for test
	helpers.HashPassword = func(pw string) (string, error) {
		assert.Equal(t, password, pw) // Optional: check input password is as expected
		return mockHash, nil
	}

	// Mock GetUserByUsername to return no rows (user does not exist)
	mock.ExpectQuery("SELECT id, username, role, tickets, created_at FROM users WHERE username = ?").
		WithArgs(username).
		WillReturnError(sql.ErrNoRows) // proper no rows simulation

	// Expect insert with the mocked hashed password
	mock.ExpectExec("INSERT INTO users").
		WithArgs(username, mockHash).
		WillReturnResult(sqlmock.NewResult(10, 1))

	id, err := repo.CreateUser(username, password)

	assert.NoError(t, err)
	assert.Equal(t, int64(10), id)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestDeleteUser_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewUserRepository(db)

	userID := int64(1)

	mock.ExpectExec("DELETE FROM users WHERE id = ?").
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.DeleteUser(userID)

	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestDeleteUser_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewUserRepository(db)

	userID := int64(999)

	mock.ExpectExec("DELETE FROM users WHERE id = ?").
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err = repo.DeleteUser(userID)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no user found")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUpdateUserRole(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewUserRepository(db)

	userID := int64(1)
	role := "admin"

	mock.ExpectExec(`UPDATE users SET role = \? WHERE id = \?`).
		WithArgs(role, userID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.UpdateUserRole(userID, role)

	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestIsAdmin_True(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewUserRepository(db)

	username := "adminuser"

	rows := sqlmock.NewRows([]string{"id", "username", "role", "tickets", "created_at"}).
		AddRow(int64(1), username, "admin", int64(0), "2025-05-17T10:00:00Z")

	mock.ExpectQuery("SELECT id, username, role, tickets, created_at FROM users WHERE username = ?").
		WithArgs(username).
		WillReturnRows(rows)

	isAdmin := repo.IsAdmin(username)

	assert.True(t, isAdmin)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestIsAdmin_False(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewUserRepository(db)

	username := "regularuser"

	rows := sqlmock.NewRows([]string{"id", "username", "role", "tickets", "created_at"}).
		AddRow(int64(2), username, "user", int64(0), "2025-05-17T10:00:00Z")

	mock.ExpectQuery("SELECT id, username, role, tickets, created_at FROM users WHERE username = ?").
		WithArgs(username).
		WillReturnRows(rows)

	isAdmin := repo.IsAdmin(username)

	assert.False(t, isAdmin)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestIsSameUser_True(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewUserRepository(db)

	userID := int64(1)
	username := "user1"

	rows := sqlmock.NewRows([]string{"id", "username", "role", "tickets", "created_at"}).
		AddRow(userID, username, "user", int64(0), "2025-05-17T10:00:00Z")

	mock.ExpectQuery("SELECT id, username, role, tickets, created_at FROM users WHERE id = ?").
		WithArgs(userID).
		WillReturnRows(rows)

	isSame := repo.IsSameUser(username, userID)

	assert.True(t, isSame)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestIsSameUser_False(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewUserRepository(db)

	userID := int64(1)
	username := "user1"

	// Return no rows (user not found)
	mock.ExpectQuery("SELECT id, username, role, tickets, created_at FROM users WHERE id = ?").
		WithArgs(userID).
		WillReturnError(sqlmock.ErrCancelled) // simulate no rows or error

	isSame := repo.IsSameUser(username, userID)

	assert.False(t, isSame)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRemoveOneTicketFromUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewUserRepository(db)

	userID := int64(1)

	mock.ExpectExec("UPDATE users SET tickets = tickets - 1 WHERE id = ?").
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.RemoveOneTicketFromUser(userID)

	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
