package main

import (
	"fmt"
)

func main() {
	// buffered channel
	messages := make(chan string, 3)

	fmt.Println("Sending messages to buffered channel")
	messages <- "First message"
	messages <- "Second message"
	messages <- "Third message"
	// messages <- "Fourth message"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	// fmt.Println(<-messages)
}
