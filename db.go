package main

import(
	"database/sql"
	"log"
	"github.com/mattn/go-sqlite3")

var db *sql.DB

func initDB()  {
	var err error
	db, err = sql.open("sqlite3", "notes.db")
	if errr != nil {
		log.Fatal(err)
	}

	createTable := '
	CREATE TABLE IF NOT EXISTS notes (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		contrnt TEXT NOT NULL
		);
		'
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}