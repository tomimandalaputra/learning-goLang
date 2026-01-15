package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActive  bool
	JoinedAt  time.Time
}

func NewEmployee(id int, firstName, lastName, position string, salary int, isActive bool) Employee {
	return Employee{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Position:  position,
		Salary:    salary,
		IsActive:  isActive,
		JoinedAt:  time.Now(),
	}
}

func main() {
	tomi := Employee{
		ID:        1,
		FirstName: "Tomi",
		LastName:  "Putra",
		Position:  "Frontend Developer",
		Salary:    1000,
		IsActive:  true,
		JoinedAt:  time.Now(),
	}
	fmt.Printf("%+v\n", tomi)

	newEmployee := NewEmployee(1, "Mandala", "Putra", "Sn Backend Developer", 1000, true)
	fmt.Printf("%+v\n", newEmployee)
}
