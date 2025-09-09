package main

import (
	"log"
	"net/http"
)

func main() {
	InitDB()
	defer DB.Close()

	http.HandleFunc("/notes", NotesHandler)

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
