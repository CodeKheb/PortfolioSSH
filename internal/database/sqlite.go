package database

import (
	"database/sql"
	"fmt"
	"time"
)

type Message struct {
	ID        int64
	Name      string
	Email     string
	Content   string
	CreatedAt time.Time
}

type Database struct {
	connection *sql.DB
}

func Open(path string) (*Database, error) {
	connection, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("db at %w", err)
	}

	database := &Database{connection: connection}

	if err := database.createTables(); err != nil {
		connection.Close()
		return nil, err
	}

	return database, nil
}

func (database *Database) createTables() error {
	const query = `
		CREATE TABLE IF NOT EXISTS messages (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			content TEXT NOT NULL,
			created_at DATETIME NOT NULL
		);
	`

	if _, err := database.connection.Exec(query); err != nil {
		return fmt.Errorf("created table at %w", err)
	}
	return nil
}

func (database *Database) SaveMessage(name, email, content string) error {
	const query = `
		INSERT INTO messages (name, email, content, created_at)
		VALUES (?, ?, ?, ?);
	`

	_, err := database.connection.Exec(
		query,
		name,
		email,
		content,
		time.Now().UTC(),
	)

	if err != nil {
		return fmt.Errorf("saved %w", err)
	}

	return nil
}

func (database *Database) Close() error {
	return database.connection.Close()
}
