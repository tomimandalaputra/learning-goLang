package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

// Database transaction
//-----------------------

// 1. User creates account
// 2. Create a wallet for the user
// 3. Want to top up the wallet for the user
// 4. You want to write a transaction log

var schema = `
CREATE TABLE IF NOT EXISTS profile (
    user_id INTEGER PRIMARY KEY REFERENCES users(user_id) ON DELETE CASCADE,
    avatar TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

type Profile struct {
	UserID    int       `json:"user_id"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	Profile        Profile   `json:"profile"`
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

	// CREATE NEW USER
	// userID, err := createUser(db, "Tomi", "tomtom@example.com", "password", "http://avatar.com/user/1")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("User created with ID:", userID)

	// GET USER BY EMAIL
	user, err := getUserByEmail(db, "tomtom@example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== Single User ===")
	result, err := formatToJSON(user)
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
	fmt.Println("Table created")
}

// Begin, Rollback or commit
func createUser(db *sql.DB, name, email, hashedPassword, avatar string) (int64, error) {
	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`)
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

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	profileStmt, err := tx.PrepareContext(ctx, `INSERT INTO profile (user_id, avatar) VALUES (?, ?)`)
	if err != nil {
		return 0, err
	}
	defer profileStmt.Close()

	_, err = profileStmt.Exec(userID, avatar)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func getUserByEmail(db *sql.DB, email string) (*User, error) {
	stmt := `SELECT u.id, u.name, u.email, u.hashed_password, u.created_at, p.avatar FROM users u
	INNER JOIN profile p ON u.id = p.user_id WHERE u.email = ?`

	row := db.QueryRow(stmt, email)
	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt, &user.Profile.Avatar)
	if err != nil {
		return nil, err
	}

	user.Profile.UserID = user.ID

	return &user, nil
}

func formatToJSON[T interface{}](data T) (string, error) {
	result, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(result), nil
}
