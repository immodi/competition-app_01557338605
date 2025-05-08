package db

import (
	"database/sql"
	"fmt"
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
	}

	for _, stmt := range schemaStatements {
		if _, err := db.Exec(stmt); err != nil {
			return fmt.Errorf("schema creation failed: %w", err)
		}
	}
	log.Println("Database schema initialized")
	return nil
}

func (db *Database) Close() error {
	log.Println("Closing database connection")
	return db.DB.Close()
}
