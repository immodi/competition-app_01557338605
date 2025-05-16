package db

import (
	"database/sql"
	"fmt"
	"immodi/submission-backend/helpers"
	"log"

	_ "modernc.org/sqlite"
)

type Database struct {
	*sql.DB
}

// NewDatabase opens a connection, runs setup, and returns a wrapped db instance.
func NewDatabase(dbPath string) (*Database, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("couldnt open database at %s: %w", dbPath, err)
	}

	// Quick connectivity test
	if pingErr := db.Ping(); pingErr != nil {
		return nil, fmt.Errorf("unable to reach database: %w", pingErr)
	}

	if err := initSchema(db); err != nil {
		return nil, err
	}

	AddDefaultAdmin(db)

	log.Printf("Connected to database: %s", dbPath)
	return &Database{DB: db}, nil
}

func initSchema(db *sql.DB) error {
	// Enforce foreign key constraints
	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return fmt.Errorf("failed to enable foreign keys: %w", err)
	}

	schemaStatements := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE NOT NULL,
			password_hash TEXT NOT NULL,
			role TEXT NOT NULL DEFAULT 'user',
			tickets INTEGER DEFAULT 999,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,

		`CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			category TEXT NOT NULL,
			date TIMESTAMP,
			venue TEXT NOT NULL,
			price REAL NOT NULL,
			image BLOB
		);`,

		`CREATE TABLE IF NOT EXISTS event_translations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_id INTEGER NOT NULL,
			language TEXT NOT NULL,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			venue TEXT NOT NULL,
			FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE
		);`,

		`CREATE TABLE IF NOT EXISTS registrations (
			user_id INTEGER NOT NULL,
			event_id INTEGER NOT NULL,
			registered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (user_id, event_id),
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
			FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE
		);`,
	}

	for _, stmt := range schemaStatements {
		if _, err := db.Exec(stmt); err != nil {
			return fmt.Errorf("schema creation failed: %w", err)
		}
	}

	log.Println("Database schema initialized")
	return nil
}

func AddDefaultAdmin(db *sql.DB) {
	password := "admin"
	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		log.Fatal("couldnt hash default admin password:", err)
	}

	insertAdmin := `
		INSERT OR IGNORE INTO users (username, password_hash, role)
		VALUES ('admin', ?, 'admin');
	`

	_, err = db.Exec(insertAdmin, hashedPassword)
	if err != nil {
		log.Fatal("Failed to insert default admin user:", err)
	}
}

func (db *Database) Close() error {
	log.Println("Closing database connection")
	return db.DB.Close()
}
