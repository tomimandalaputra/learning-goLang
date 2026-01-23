package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"learning-go/section12/6-repository/repository"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbName := "users_database.db"
	db, err := connectToDB(dbName)
	checkError(err)
	defer db.Close()

	fmt.Println("Database connection established")

	repo := repository.NewSQLUserRepository(db)

	// === INSERT DATA USER ===
	// insertUser(repo)

	// === PRINT ALL USER ===
	printUsers(repo)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func connectToDB(dbName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func insertUser(repo repository.UserRepository) {
	user, err := repo.CreateUser("test", "test@mail.com", "pass", "avatartest")
	checkError(err)
	fmt.Println("User created: ", user)
}

func printUsers(repo repository.UserRepository) {
	users, err := repo.GetUsers()
	checkError(err)

	json, err := formatToJSON(users)
	checkError(err)
	fmt.Println(json)
}

func formatToJSON[T interface{}](data T) (string, error) {
	result, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(result), nil
}
