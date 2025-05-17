package tests

import (
	"database/sql"
	"immodi/submission-backend/repos"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllEvents(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewEventRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name", "description", "category", "date", "venue", "price"}).
		AddRow(int64(1), "Event1", "Desc1", "Cat1", "2025-01-01", "Venue1", 10.0).
		AddRow(int64(2), "Event2", "Desc2", "Cat2", "2025-02-02", "Venue2", 20.0)

	mock.ExpectQuery("SELECT id, name, description, category, date, venue, price FROM events").
		WillReturnRows(rows)

	events, err := repo.GetAllEvents()
	assert.NoError(t, err)
	assert.Len(t, events, 2)
	assert.Equal(t, "Event1", events[0].Name)
	assert.Equal(t, "Event2", events[1].Name)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetEventById_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewEventRepository(db)

	eventID := int64(1)

	// Mock event row
	eventRows := sqlmock.NewRows([]string{"id", "name", "description", "category", "date", "venue", "price", "image"}).
		AddRow(eventID, "Event1", "Desc1", "Cat1", "2025-01-01", "Venue1", 10.0, []byte{1, 2, 3})

	mock.ExpectQuery("SELECT id, name, description, category, date, venue, price, image FROM events WHERE id = ?").
		WithArgs(eventID).
		WillReturnRows(eventRows)

	// Mock translations
	transRows := sqlmock.NewRows([]string{"language", "name", "description", "venue"}).
		AddRow("en", "Event1 EN", "Desc EN", "Venue EN").
		AddRow("fr", "Event1 FR", "Desc FR", "Venue FR")

	mock.ExpectQuery("SELECT language, name, description, venue FROM event_translations WHERE event_id = ?").
		WithArgs(eventID).
		WillReturnRows(transRows)

	event, err := repo.GetEventById(eventID)
	assert.NoError(t, err)
	assert.NotNil(t, event)
	assert.Equal(t, eventID, event.ID)
	assert.Len(t, event.Translations, 2)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetEventById_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewEventRepository(db)

	eventID := int64(999)

	mock.ExpectQuery("SELECT id, name, description, category, date, venue, price, image FROM events WHERE id = ?").
		WithArgs(eventID).
		WillReturnError(sql.ErrNoRows)

	event, err := repo.GetEventById(eventID)
	assert.NoError(t, err)
	assert.Nil(t, event)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetEventsByCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewEventRepository(db)

	category := "Cat1"

	rows := sqlmock.NewRows([]string{"id", "name", "description", "category", "date", "venue", "price"}).
		AddRow(int64(1), "Event1", "Desc1", category, "2025-01-01", "Venue1", 10.0)

	mock.ExpectQuery("SELECT id, name, description, category, date, venue, price FROM events WHERE category = ?").
		WithArgs(category).
		WillReturnRows(rows)

	events, err := repo.GetEventsByCategory(category)
	assert.NoError(t, err)
	assert.Len(t, events, 1)
	assert.Equal(t, category, events[0].Category)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestCreateEvent(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewEventRepository(db)

	name := "Event1"
	description := "Desc1"
	category := "Cat1"
	date := "2025-01-01"
	venue := "Venue1"
	price := 10.0
	image := []byte{1, 2, 3}

	eventTranslations := []repos.EventTranslation{
		{Language: "en", Name: "Name EN", Description: "Desc EN", Venue: "Venue EN"},
	}

	mock.ExpectExec("INSERT INTO events").
		WithArgs(name, description, category, date, venue, price, image).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT INTO event_translations").
		WithArgs(int64(1), eventTranslations[0].Language, eventTranslations[0].Name, eventTranslations[0].Description, eventTranslations[0].Venue).
		WillReturnResult(sqlmock.NewResult(1, 1))

	id, err := repo.CreateEvent(name, description, category, date, venue, price, image, eventTranslations)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), id)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUpdateEvent(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewEventRepository(db)

	id := int64(1)
	name := "Updated"
	description := "Updated Desc"
	category := "Updated Cat"
	date := "2025-02-02"
	venue := "Updated Venue"
	price := 20.0
	image := []byte{4, 5, 6}

	eventTranslations := []repos.EventTranslation{
		{Language: "en", Name: "Updated EN", Description: "Desc EN", Venue: "Venue EN"},
	}

	mock.ExpectExec("UPDATE events").
		WithArgs(name, description, category, date, venue, price, image, id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectExec("DELETE FROM event_translations WHERE event_id = ?").
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectExec("INSERT INTO event_translations").
		WithArgs(id, eventTranslations[0].Language, eventTranslations[0].Name, eventTranslations[0].Description, eventTranslations[0].Venue).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.UpdateEvent(id, name, description, category, date, venue, price, image, eventTranslations)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestDeleteEvent(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewEventRepository(db)

	id := int64(1)

	mock.ExpectExec("DELETE FROM events WHERE id = ?").
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.DeleteEvent(id)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetUpcomingEvents(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewEventRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name", "description", "category", "date", "venue", "price", "image"}).
		AddRow(int64(1), "Upcoming Event", "Desc", "Cat", "2025-06-01", "Venue", 30.0, []byte{1, 2})

	mock.ExpectQuery("SELECT id, name, description, category, date, venue, price, image FROM events WHERE date >= datetime\\('now'\\) ORDER BY date ASC").
		WillReturnRows(rows)

	events, err := repo.GetUpcomingEvents()
	assert.NoError(t, err)
	assert.Len(t, events, 1)
	assert.Equal(t, "Upcoming Event", events[0].Name)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestSearchEvents(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewEventRepository(db)

	keyword := "party"
	searchTerm := "%" + keyword + "%"

	rows := sqlmock.NewRows([]string{"id", "name", "description", "category", "date", "venue", "price"}).
		AddRow(int64(1), "Party Event", "Desc", "Fun", "2025-07-07", "Club", 50.0)

	mock.ExpectQuery("SELECT id, name, description, category, date, venue, price FROM events WHERE name LIKE ?").
		WithArgs(searchTerm, searchTerm, searchTerm).
		WillReturnRows(rows)

	events, err := repo.SearchEvents(keyword)
	assert.NoError(t, err)
	assert.Len(t, events, 1)
	assert.Equal(t, "Party Event", events[0].Name)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRegisterUserToEvent(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewEventRepository(db)

	userID := int64(1)
	eventID := int64(2)

	mock.ExpectExec("INSERT INTO registrations \\(user_id, event_id\\) VALUES \\(\\?, \\?\\)").
		WithArgs(userID, eventID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.RegisterUserToEvent(userID, eventID)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetEventsForUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repos.NewEventRepository(db)

	userID := int64(1)

	rows := sqlmock.NewRows([]string{"id", "name", "description", "category", "date", "venue", "price", "image"}).
		AddRow(int64(1), "User Event", "Desc", "Cat", "2025-08-01", "Venue", 40.0, []byte{1, 2})

	mock.ExpectQuery("SELECT e.id, e.name, e.description, e.category, e.date, e.venue, e.price, e.image FROM events e JOIN registrations r ON e.id = r.event_id WHERE r.user_id = ?").
		WithArgs(userID).
		WillReturnRows(rows)

	events, err := repo.GetEventsForUser(userID)
	assert.NoError(t, err)
	assert.Len(t, events, 1)
	assert.Equal(t, "User Event", events[0].Name)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
