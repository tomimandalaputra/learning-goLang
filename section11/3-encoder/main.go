package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Position string `json:"position"`
	IsActive bool   `json:"is_active"`
}

func main() {
	e := employee{
		Name:     "Putra",
		Email:    "example@mail.com",
		IsActive: true,
	}

	// first option
	// enc := json.NewEncoder(os.Stdout)
	// if err := enc.Encode(&e); err != nil {
	// 	log.Fatal(err)
	// }

	// second option
	buff := new(bytes.Buffer)
	enc := json.NewEncoder(buff)
	if err := enc.Encode(&e); err != nil {
		log.Fatal(err)
	}

	fmt.Println(buff.String())
}
