package db

import (
	"database/sql"
	"time"

	"github.com/JasperMunene/notes-api/models"
	_ "github.com/lib/pq"
)

// New opens a PostgreSQL database connection using the given DSN.
func New(dsn string) (*sql.DB, error) {
	return sql.Open("postgres", dsn)
}

// CreateNote inserts a new note into the PostgreSQL database and sets its ID and CreatedAt.
func CreateNote(db *sql.DB, n *models.Note) error {
	n.CreatedAt = time.Now()
	return db.QueryRow(
		"INSERT INTO notes (title, content, created) VALUES ($1, $2, $3) RETURNING id",
		n.Title, n.Content, n.CreatedAt,
	).Scan(&n.ID)
}
