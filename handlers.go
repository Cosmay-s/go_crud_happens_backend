package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func NotesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		handleCreateNote(w, r)

	case http.MethodGet:
		handleGetAllNotes(w, r)

	default:
		http.Error(w, "Method error", http.StatusMethodNotAllowed)
	}
}

func NoteByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/notes/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "ID error", http.StatusBadRequest)
		return
	}
	switch r.Method {

	case http.MethodGet:
		handleGetNoteByID(w, r, id)

	case http.MethodPut:
		handleUpdateNote(w, r, id)

	case http.MethodDelete:
		handleDeleteNote(w, r, id)

	default:
		http.Error(w, "getbyid method error", http.StatusMethodNotAllowed)
	}
}

func handleCreateNote(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type error", http.StatusUnsupportedMediaType)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Read body error", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var note Note
	err = json.Unmarshal(body, &note)
	if err != nil {
		http.Error(w, "json error", http.StatusBadRequest)
		return
	}

	if note.Title == "" || note.Content == "" {
		http.Error(w, "title or content error", http.StatusBadRequest)
		return
	}

	err = CreateNote(&note)
	if err != nil {
		http.Error(w, "CreateNote error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

func CreateNote(note *Note) error {
	result, err := DB.Exec("INSERT INTO notes (title, content) VALUES (?, ?)", note.Title, note.Content)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	note.ID = int(id)
	return nil
}

func handleGetAllNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := GetAllNotes()
	if err != nil {
		http.Error(w, "get error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func GetAllNotes() ([]Note, error) {
	rows, err := DB.Query("SELECT id, title, content FROM notes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var n Note
		err := rows.Scan(&n.ID, &n.Title, &n.Content)
		if err != nil {
			return nil, err
		}
		notes = append(notes, n)
	}
	return notes, rows.Err()
}

func handleGetNoteByID(w http.ResponseWriter, r *http.Request, id int) {
	note, err := GetNoteByID(id)
	if err != nil {
		http.Error(w, "Note ID not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func GetNoteByID(id int) (*Note, error) {
	row := DB.QueryRow("SELECT id, title, content FROM notes WHERE id = ?", id)
	var n Note
	err := row.Scan(&n.ID, &n.Title, &n.Content)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

func handleUpdateNote(w http.ResponseWriter, r *http.Request, id int) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type error", http.StatusUnsupportedMediaType)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Read body error", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var note Note
	err = json.Unmarshal(body, &note)
	if err != nil {
		http.Error(w, "json error", http.StatusBadRequest)
		return
	}

	if note.Title == "" || note.Content == "" {
		http.Error(w, "title or content error", http.StatusBadRequest)
		return
	}
	note.ID = id

	err = UpdateNote(&note)

	if err != nil {
		http.Error(w, "Update error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

func UpdateNote(note *Note) error {
	result, err := DB.Exec("UPDATE notes SET title = ?, content = ? WHERE id = ?", note.Title, note.Content, note.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func handleDeleteNote(w http.ResponseWriter, r *http.Request, id int) {
	err := DeleteNote(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Note not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to delete note", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteNote(id int) error {
	result, err := DB.Exec("DELETE FROM notes WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
