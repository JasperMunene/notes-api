package main

import (
	"fmt"
	"github.com/JasperMunene/notes-api/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/notes", handlers.NotesHandler)

	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
