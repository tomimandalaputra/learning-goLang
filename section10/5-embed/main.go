package main

import (
	"embed"
	"fmt"
	"log"
)

var name = "Tomi Mandala Putra"

//go:embed public
var public embed.FS

func main() {
	data, err := public.ReadFile("public/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	fmt.Println(name)
}
