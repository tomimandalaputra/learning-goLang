package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	data := "Hello World, welcome to the wonderful world of Go!"

	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Result of base64:", encoded)

	decodedStr, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(decodedStr))
}
