package main

import (
	"log"
	"net/http"
)

func main() {
	initDB()

	http.HandleFunc("/notes", notesHandler)
	http.HandleFunc("/notes", noteByIDHandler)

	log.Println("Сервер запущен")
}
