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

func UpdateNote(db *sql.DB, id string, n *models.Note) error {
	_, err := db.Exec(
		"UPDATE notes SET title = $1, content = $2 WHERE id = $3",
		n.Title, n.Content, id,
	)
	return err
}

func DeleteNote(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM notes WHERE id = $1", id)
	return err
}
