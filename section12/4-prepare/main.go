package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

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

	// Example: Create user with prepared statement
	// _, err = createUserWithPrepared(db, "ronaldo", "ronaldo@example.com", "password123")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Example: Create user with prepared statement and context
	ctx := context.Background()
	_, err = createUserWithContext(ctx, db, "Ronal", "ronal@example.com", "password123")
	if err != nil {
		log.Fatal(err)
	}
}

func createUserWithPrepared(db *sql.DB, name, email, hashedPassword string) (int64, error) {
	stmt, err := db.Prepare(`INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	hp, err := bcrypt.GenerateFromPassword([]byte(hashedPassword), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(name, email, string(hp))
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func createUserWithContext(ctx context.Context, db *sql.DB, name, email, hashedPassword string) (int64, error) {
	stmt, err := db.Prepare(`INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	hp, err := bcrypt.GenerateFromPassword([]byte(hashedPassword), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	result, err := stmt.ExecContext(ctx, name, email, string(hp))
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}
