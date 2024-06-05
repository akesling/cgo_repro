package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		log.Fatalf("Opening sqlite db (db.sqlite3) failed: %v", err)
	}

	const create string = "CREATE TABLE IF NOT EXISTS test (message TEXT);"
	if _, err := db.Exec(create); err != nil {
		log.Fatalf("Creating test table in (db.sqlite3) failed: %v", err)
	}

	const insert string = "INSERT INTO test VALUES(?);"
	if _, err := db.Exec(insert, "Hello SQLite!"); err != nil {
		log.Fatalf("Inserting into test table failed: %v", err)
	}

	rows, err := db.Query("SELECT * FROM test;")
	if err != nil {
		log.Fatalf("Selecting from test table failed: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var testText string
		if err := rows.Scan(&testText); err != nil {
			log.Fatalf("Reading test text from column failed: %v", err)
		}
		fmt.Printf("Read text from DB: %s\n", testText)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("An error occurred iterating through rows: %v", err)
	}
}
