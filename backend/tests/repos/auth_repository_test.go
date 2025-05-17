package tests

import (
	"immodi/submission-backend/repos"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAuthUserByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewAuthRepository(db)

	username := "testuser"
	expectedUser := &repos.AuthUser{
		ID:           1,
		Username:     "testuser",
		Role:         "admin",
		CreatedAt:    "2024-05-17T10:00:00Z",
		PasswordHash: "hashedpassword",
	}

	rows := sqlmock.NewRows([]string{"id", "username", "role", "created_at", "password_hash"}).
		AddRow(expectedUser.ID, expectedUser.Username, expectedUser.Role, expectedUser.CreatedAt, expectedUser.PasswordHash)

	mock.ExpectQuery("SELECT id, username, role, created_at, password_hash FROM users WHERE username = ?").
		WithArgs(username).
		WillReturnRows(rows)

	user, err := repo.GetAuthUserByUsername(username)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser, user)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
