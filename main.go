package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/JasperMunene/notes-api/db"
	"github.com/JasperMunene/notes-api/handlers"
)

func main() {
	// Read DATABASE_URL
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL must be set")
	}

	// Open the database
	sqlDB, err := db.New(dsn)
	if err != nil {
		log.Fatalf("db.New: %v", err)
	}
	defer sqlDB.Close()

	// Give the handlers access to *sql.DB
	h := handlers.New(sqlDB)

	http.HandleFunc("/notes", h.NotesHandler)

	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
