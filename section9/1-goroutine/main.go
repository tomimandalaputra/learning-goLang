package main

import (
	"fmt"
	"time"
)

func sayHello(message string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Printf("sayHello %s", message)
}

func main() {
	fmt.Println("Mai() Go Routine")
	go sayHello("Lorem ipsum", time.Second)
	fmt.Println("Last Message from Main()")
	time.Sleep(2 * time.Second)
}
