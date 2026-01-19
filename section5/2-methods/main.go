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

/*
Method dengan receiver value
Receiver value tidak akan mengubah data struct
*/
func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}

/*
Method dengan receiver pointer
Receiver pointer akan mengubah data struct
*/
func (e *Employee) Deactivated() {
	e.IsActive = false
}

/*
Method SetJoinedAt digunakan untuk mengubah data struct
pada field JoinedAt dengan parameter time.Time
*/
func (e *Employee) SetJoinedAt(joinedAt time.Time) {
	e.JoinedAt = joinedAt
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

	// fmt.Println(tomi.FullName())

	// tomi.Deactivated()
	// fmt.Printf("%+v\n", tomi)

	tomi.SetJoinedAt(time.Now().Add(7 * time.Hour))
	fmt.Printf("%+v\n", tomi)
}
