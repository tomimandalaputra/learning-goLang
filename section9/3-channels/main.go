package main

import (
	"fmt"
	"time"
)

type UserDetail struct {
	name string
}

func main() {
	// channel
	messages := make(chan string)

	user := make(chan UserDetail)

	go func() {
		fmt.Println("Sending a message to message channel")
		messages <- "Hello from message channel"
	}()

	// go func() {
	// 	fmt.Println("Sending a message to message channel")
	// 	messages <- "Hello from message channel"
	// }()

	go func() {
		fmt.Println("Sending a message to user channel")
		user <- UserDetail{
			name: "Tomi Mandala Putra",
		}
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("About to get message from channel")

	msg := <-messages
	fmt.Println(msg)

	// msg = <-messages
	// fmt.Println(msg)

	u := <-user
	fmt.Printf("%+v\n", u)
}
