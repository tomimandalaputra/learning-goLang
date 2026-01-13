package main

import "fmt"

func main() {
	var greeting string        // zero-value is an empty string "" -> this is declaration variable
	greeting = "Hello, world!" // this is initiation value to variable greeting
	fmt.Println(greeting)

	var count int
	count = 8
	fmt.Println(count)

	var isRunning bool
	isRunning = true
	fmt.Println(isRunning)

	var firstName, lastName string
	firstName = "Tomi"
	lastName = "Putra"
	fmt.Println(firstName, lastName)

	email := "example@mail.com"
	fmt.Println(email)

	age := 25
	fmt.Println(age)

	var year = 2026
	fmt.Println(year)
}
