package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func New() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn != "" {
		return sql.Open("postgres", dsn)
	}
	// fallback to SQLite
	return sql.Open("sqlite3", "file:notes.db?_foreign_keys=on")
}
