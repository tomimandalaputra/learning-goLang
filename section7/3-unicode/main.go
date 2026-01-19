package main

import (
	"fmt"
	"unicode"
)

func main() {
	// words := "水漢机"
	words := []rune{'水', '漢', '机'}

	for _, v := range words {
		fmt.Printf("%s, %t\n", string(v), unicode.IsLetter(v))
	}
}
