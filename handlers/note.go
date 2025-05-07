// Package handlers provides HTTP handlers for the notes API.
package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/JasperMunene/notes-api/db"
	"github.com/JasperMunene/notes-api/models"
)

// Handler wraps dependencies for HTTP handlers.
type Handler struct {
	db *sql.DB
}

// New returns a Handler ready to use.
func New(dbConn *sql.DB) *Handler {
	return &Handler{db: dbConn}
}

// NotesHandler dispatches to the appropriate method based on HTTP verb.
func (h *Handler) NotesHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	switch r.Method {
	case http.MethodGet:
		h.getNotes(w, r)
	case http.MethodPost:
		h.createNote(w, r)
	case http.MethodPut:
		if id == "" {
			h.respondError(w, http.StatusBadRequest, "missing note id")
			return
		}
		h.updateNote(w, r, id)
	case http.MethodDelete:
		if id == "" {
			h.respondError(w, http.StatusBadRequest, "missing note id")
			return
		}
		h.deleteNote(w, r, id)
	default:
		h.respondError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

// GET /notes
func (h *Handler) getNotes(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.Query(
		"SELECT id, title, content, created FROM notes ORDER BY created DESC",
	)
	if err != nil {
		log.Printf("getNotes query error: %v", err)
		h.respondError(w, http.StatusInternalServerError, fmt.Sprintf("query error: %v", err))
		return
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var n models.Note
		if err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt); err != nil {
			log.Printf("getNotes scan error: %v", err)
			h.respondError(w, http.StatusInternalServerError, fmt.Sprintf("scan error: %v", err))
			return
		}
		notes = append(notes, n)
	}

	h.respondJSON(w, http.StatusOK, notes)
}

// POST /notes
func (h *Handler) createNote(w http.ResponseWriter, r *http.Request) {
	var n models.Note
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		h.respondError(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	if err := db.CreateNote(h.db, &n); err != nil {
		log.Printf("createNote error: %v", err)
		h.respondError(w, http.StatusInternalServerError, fmt.Sprintf("could not create note: %v", err))
		return
	}

	h.respondJSON(w, http.StatusCreated, n)
}

// PUT /notes
func (h *Handler) updateNote(w http.ResponseWriter, r *http.Request, id string) {
	var n models.Note
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		h.respondError(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	if err := db.UpdateNote(h.db, id, &n); err != nil {
		log.Printf("updateNote error: %v", err)
		h.respondError(w, http.StatusInternalServerError, fmt.Sprintf("could not update note: %v", err))
		return
	}

	h.respondJSON(w, http.StatusOK, map[string]string{"message": "note updated"})
}

// DELETE /notes
func (h *Handler) deleteNote(w http.ResponseWriter, r *http.Request, id string) {
	if err := db.DeleteNote(h.db, id); err != nil {
		log.Printf("deleteNote error: %v", err)
		h.respondError(w, http.StatusInternalServerError, fmt.Sprintf("could not delete note: %v", err))
		return
	}

	h.respondJSON(w, http.StatusOK, map[string]string{"message": "note deleted"})
}

// respondJSON writes a JSON response with the given status code
func (h *Handler) respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// respondError writes an error message as JSON
func (h *Handler) respondError(w http.ResponseWriter, status int, message string) {
	h.respondJSON(w, status, map[string]string{"error": message})
}
