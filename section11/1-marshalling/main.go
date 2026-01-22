package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Employee struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Position string `json:"position"`
	IsActive bool   `json:"is_active"`
}

func main() {
	employee := Employee{
		Name:     "Tomi",
		Email:    "exmaple@mail.com",
		Position: "Software Developer",
		IsActive: true,
	}
	// result, err := json.Marshal(employee)
	result, err := json.MarshalIndent(employee, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result))
}
