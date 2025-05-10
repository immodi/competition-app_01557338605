package repos

import (
	"database/sql"
	"fmt"
)

type Event struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Date        string  `json:"date"`
	Venue       string  `json:"venue"`
	Price       float64 `json:"price"`
	Image       []byte  `json:"image,omitempty"`
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

	var events []Event
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
	return &e, nil
}

func (r *EventRepository) GetEventsByCategory(category string) ([]Event, error) {
	rows, err := r.db.Query("SELECT id, name, description, category, date, venue, price, image FROM events WHERE category = ?", category)
	if err != nil {
		return nil, fmt.Errorf("fetching events by category failed: %w", err)
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Category, &e.Date, &e.Venue, &e.Price, &e.Image); err != nil {
			return nil, fmt.Errorf("error scanning event by category: %w", err)
		}
		events = append(events, e)
	}

	return events, rows.Err()
}

func (r *EventRepository) CreateEvent(name, description, category, date, venue string, price float64, image []byte) (int64, error) {
	result, err := r.db.Exec(
		`INSERT INTO events (name, description, category, date, venue, price, image) 
		 VALUES (?, ?, ?, ?, ?, ?, ?)`,
		name, description, category, date, venue, price, image,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to create event: %w", err)
	}
	return result.LastInsertId()
}

func (r *EventRepository) UpdateEvent(id int64, name, description, category, date, venue string, price float64, image []byte) error {
	_, err := r.db.Exec(
		`UPDATE events 
		 SET name = ?, description = ?, category = ?, date = ?, venue = ?, price = ?, image = ? 
		 WHERE id = ?`,
		name, description, category, date, venue, price, image, id,
	)
	if err != nil {
		return fmt.Errorf("failed to update event id %d: %w", id, err)
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

	var events []Event
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

	var events []Event
	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Category, &e.Date, &e.Venue, &e.Price, &e.Image); err != nil {
			return nil, fmt.Errorf("error scanning search result: %w", err)
		}
		events = append(events, e)
	}
	return events, rows.Err()
}
