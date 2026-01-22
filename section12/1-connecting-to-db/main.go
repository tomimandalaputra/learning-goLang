package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    hashed_password BLOB NOT NULL, -- Storing as BLOB for byte slice
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

func main() {
	dbName := "data.db"

	_ = os.Remove(dbName) // very important you do not do write do this in production!

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		fmt.Println("Closing database connection")
		if err := db.Close(); err != nil {
			log.Printf("Error clossing database connection: %v", err)
		}
	}()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connection established")

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table was created")
}
