package repos

import (
	"database/sql"
	"fmt"
)

type Event struct {
	ID           int64              `json:"id"`
	Name         string             `json:"name"`
	Description  string             `json:"description"`
	Category     string             `json:"category"`
	Date         string             `json:"date"`
	Venue        string             `json:"venue"`
	Price        float64            `json:"price"`
	Image        []byte             `json:"image,omitempty"`
	Translations []EventTranslation `json:"translations"`
}

type EventTranslation struct {
	Language    string `json:"language"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Venue       string `json:"venue"`
}

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) GetAllEvents() ([]Event, error) {
	rows, err := r.db.Query("SELECT id, name, description, category, date, venue, price, image FROM events")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch events: %w", err)
	}
	defer rows.Close()

	events := []Event{}
	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Category, &e.Date, &e.Venue, &e.Price, &e.Image); err != nil {
			return nil, fmt.Errorf("error scanning event row: %w", err)
		}
		events = append(events, e)
	}

	return events, rows.Err()
}

func (r *EventRepository) GetEventById(id int64) (*Event, error) {
	var e Event
	query := "SELECT id, name, description, category, date, venue, price, image FROM events WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&e.ID, &e.Name, &e.Description, &e.Category, &e.Date, &e.Venue, &e.Price, &e.Image)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get event by id %d: %w", id, err)
	}

	e.Translations, err = r.GetEventTranslations(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get event translations by id %d: %w", id, err)
	}
	return &e, nil
}

func (r *EventRepository) GetEventsByCategory(category string) ([]Event, error) {
	rows, err := r.db.Query("SELECT id, name, description, category, date, venue, price, image FROM events WHERE category = ?", category)
	if err != nil {
		return nil, fmt.Errorf("fetching events by category failed: %w", err)
	}
	defer rows.Close()

	events := []Event{}
	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Category, &e.Date, &e.Venue, &e.Price, &e.Image); err != nil {
			return nil, fmt.Errorf("error scanning event by category: %w", err)
		}
		events = append(events, e)
	}

	return events, rows.Err()
}

func (r *EventRepository) CreateEvent(name, description, category, date, venue string, price float64, image []byte, eventTranslations []EventTranslation) (int64, error) {
	result, err := r.db.Exec(
		`INSERT INTO events (name, description, category, date, venue, price, image) 
		 VALUES (?, ?, ?, ?, ?, ?, ?)`,
		name, description, category, date, venue, price, image,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to create event: %w", err)
	}

	eventId, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}

	for _, et := range eventTranslations {
		_, err := r.db.Exec(
			`INSERT INTO event_translations (event_id, language, name, description, venue) 
			 VALUES (?, ?, ?, ?, ?)`,
			eventId, et.Language, et.Name, et.Description, et.Venue,
		)
		if err != nil {
			return 0, fmt.Errorf("failed to create event translation: %w", err)
		}
	}
	return result.LastInsertId()
}

func (r *EventRepository) UpdateEvent(id int64, name, description, category, date, venue string, price float64, image []byte, eventTranslations []EventTranslation) error {
	_, err := r.db.Exec(
		`UPDATE events 
		 SET name = ?, description = ?, category = ?, date = ?, venue = ?, price = ?, image = ? 
		 WHERE id = ?`,
		name, description, category, date, venue, price, image, id,
	)
	if err != nil {
		return fmt.Errorf("failed to update event id %d: %w", id, err)
	}

	_, err = r.db.Exec(`DELETE FROM event_translations WHERE event_id = ?`, id)
	if err != nil {
		return fmt.Errorf("failed to delete existing event translations: %w", err)
	}

	for _, et := range eventTranslations {
		_, err := r.db.Exec(
			`INSERT INTO event_translations (event_id, language, name, description, venue) 
         VALUES (?, ?, ?, ?, ?)`,
			id, et.Language, et.Name, et.Description, et.Venue,
		)
		if err != nil {
			return fmt.Errorf("failed to insert event translation: %w", err)
		}
	}

	return nil
}

func (r *EventRepository) DeleteEvent(id int64) error {
	_, err := r.db.Exec("DELETE FROM events WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete event id %d: %w", id, err)
	}
	return nil
}

func (r *EventRepository) GetUpcomingEvents() ([]Event, error) {
	rows, err := r.db.Query(
		`SELECT id, name, description, category, date, venue, price, image 
		 FROM events 
		 WHERE date >= datetime('now') 
		 ORDER BY date ASC`,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch upcoming events: %w", err)
	}
	defer rows.Close()

	events := []Event{}
	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Category, &e.Date, &e.Venue, &e.Price, &e.Image); err != nil {
			return nil, fmt.Errorf("error scanning upcoming event: %w", err)
		}
		events = append(events, e)
	}

	return events, rows.Err()
}

func (r *EventRepository) SearchEvents(keyword string) ([]Event, error) {
	searchTerm := "%" + keyword + "%"
	rows, err := r.db.Query(
		`SELECT id, name, description, category, date, venue, price, image 
		 FROM events 
		 WHERE name LIKE ? OR description LIKE ? OR venue LIKE ?`,
		searchTerm, searchTerm, searchTerm,
	)
	if err != nil {
		return nil, fmt.Errorf("search query failed: %w", err)
	}
	defer rows.Close()

	events := []Event{}
	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Category, &e.Date, &e.Venue, &e.Price, &e.Image); err != nil {
			return nil, fmt.Errorf("error scanning search result: %w", err)
		}
		events = append(events, e)
	}
	return events, rows.Err()
}

func (r *EventRepository) RegisterUserToEvent(userID, eventID int64) error {
	_, err := r.db.Exec(`
		INSERT INTO registrations (user_id, event_id)
		VALUES (?, ?)
	`, userID, eventID)
	if err != nil {
		return fmt.Errorf("failed to register user %d to event %d: %w", userID, eventID, err)
	}
	return nil
}

func (r *EventRepository) GetEventsForUser(userID int64) ([]Event, error) {
	rows, err := r.db.Query(
		`SELECT e.id, e.name, e.description, e.category, e.date, e.venue, e.price, e.image
		 FROM events e
		 JOIN registrations r ON e.id = r.event_id
		 WHERE r.user_id = ?`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []Event{}
	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Category, &e.Date, &e.Venue, &e.Price, &e.Image); err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func (r *EventRepository) GetEventTranslations(eventId int64) ([]EventTranslation, error) {
	rows, err := r.db.Query("SELECT language, name, description, venue FROM event_translations WHERE event_id = ?", eventId)
	if err != nil {
		return nil, fmt.Errorf("fetching events by category failed: %w", err)
	}
	defer rows.Close()

	eventTranslations := []EventTranslation{}
	for rows.Next() {
		var e EventTranslation
		if err := rows.Scan(&e.Language, &e.Name, &e.Description, &e.Venue); err != nil {
			return nil, fmt.Errorf("error scanning event by category: %w", err)
		}
		eventTranslations = append(eventTranslations, e)
	}

	return eventTranslations, rows.Err()
}
