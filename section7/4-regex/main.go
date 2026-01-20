package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {

	sentences := "Hello world! Welcome to Go"

	regexGo, err := regexp.Compile("Go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Text '%s', matcher 'Go': %t\n", sentences, regexGo.MatchString(sentences))

	text := "Products code: A123, B455, C596, A243"
	regexProductCode := regexp.MustCompile(`A\d+`)
	findFirstProduct := regexProductCode.FindString(text)
	fmt.Println(findFirstProduct)

	findAllProduct := regexProductCode.FindAllString(text, -1)
	fmt.Println(findAllProduct)
}
