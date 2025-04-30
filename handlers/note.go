package handlers

import (
	"fmt"
	"net/http"
)

func NotesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "List all notes")
	case http.MethodPost:
		fmt.Fprintln(w, "Create a new note")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
