package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "abc"
	clone := strings.Clone(word)
	fmt.Println(word)
	fmt.Println(clone)

	b := strings.Builder{}
	// b.Write([]byte("Hello, World!"))
	b.WriteString("Hello, World!")
	fmt.Println(b.String())

	fullName := "  Tomi Mandala Putra  "
	fmt.Println(strings.ToLower(fullName))
	fmt.Println(strings.ToUpper(fullName))
	fmt.Println(strings.Title(fullName))
	fmt.Println(strings.TrimSpace(fullName))

	email := "test@mail.com"
	fmt.Println(strings.HasSuffix(email, "mail.com"))
	fmt.Println(strings.HasPrefix(email, "test"))
	fmt.Println(strings.Replace(email, "test", "tomi", 1))

	split := strings.Split(email, "@")
	username, domain := split[0], split[1]
	fmt.Printf("Username: %s, domain: %s\n", username, domain)
}
