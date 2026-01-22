package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
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

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
}

func main() {
	dbName := "users_database.db"

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

	// createTable(db)

	// Example: Get single user by email
	tomi, err := getUserByEmail(db, "tomi@example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== Single User ===")
	result, err := formatToJSON(tomi)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

	// Example: Get all users
	users, err := getAllUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n=== All Users ===")
	result, err = formatToJSON(users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func createTable(db *sql.DB) {
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table was created")
}

func createUsers(db *sql.DB, name, email, hashedPassword string) (int64, error) {
	stmt := `INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`

	hp, err := bcrypt.GenerateFromPassword([]byte(hashedPassword), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	result, err := db.Exec(stmt, name, email, string(hp))
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func getUserByEmail(db *sql.DB, email string) (*User, error) {
	stmt := `SELECT * FROM users WHERE email = ?`

	row := db.QueryRow(stmt, email)

	var u User
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.HashedPassword, &u.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func getAllUsers(db *sql.DB) ([]User, error) {
	stmt := `SELECT * FROM users`

	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// define variable users
	var users []User

	for rows.Next() {
		var u User

		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.HashedPassword, &u.CreatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func formatToJSON[T interface{}](data T) (string, error) {
	result, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(result), nil
}
