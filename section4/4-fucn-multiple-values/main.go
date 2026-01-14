package main

import (
	"errors"
	"fmt"
	"strings"
)

func divide(numerator, denominator float64) (float64, error) {
	if denominator == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, nil
}

func splitName(fullName string) (firstName, lastName string) {
	parts := strings.Split(fullName, " ")
	firstName = parts[0]
	lastName = parts[1]

	/**
	* retrun dibuat tanpa nilai eksplisit
	* Karena return value sudah dinamai -> (firstName, lastName string)
	* Go otomatis mengembalikan firstName dan lastName
	 */
	return
}

func main() {
	if value, err := divide(10, 2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	fmt.Println(splitName("Tomi Mandala Putra"))
}
