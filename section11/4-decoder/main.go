package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type employee struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Position string `json:"position"`
	IsActive bool   `json:"is_active"`
}

var jsonPayload = `{"name":"Putra","email":"example@mail.com","position":"","is_active":true}`

func main() {
	var e employee

	dec := json.NewDecoder(strings.NewReader(jsonPayload))
	if err := dec.Decode(&e); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", e)
}
