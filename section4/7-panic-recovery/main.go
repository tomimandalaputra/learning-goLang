package main

import "fmt"

func mightPanic(shouldPanic bool) {
	if shouldPanic {
		panic("something went wrong in mightPanic")
	}

	fmt.Println("This func executed whitout a panic")
}

func recoverable() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Revocered from panic:", r)
		}
	}()

	mightPanic(true)
}

func main() {
	recoverable()
}
