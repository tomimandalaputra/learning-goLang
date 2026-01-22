package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	dir := "Downloads/static/images"
	if err := os.MkdirAll(filepath.Clean(dir), 0755); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully created directory")

	time.Sleep(3 * time.Second)

	if err := os.RemoveAll("Downloads"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully removed directory")
}
