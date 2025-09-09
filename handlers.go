package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func NotesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handleCreateNote(w, r)
	default:
		http.Error(w, "Method error", http.StatusMethodNotAllowed)
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
