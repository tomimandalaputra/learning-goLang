package main

import "fmt"

const HOST string = "localhost"

func main() {
	fmt.Println(HOST)

	const appName string = "Go"
	fmt.Println(appName)

	const pi float64 = 3.1415926
	fmt.Println(pi)

	const rate float32 = 8.9
	fmt.Println(rate)
}
