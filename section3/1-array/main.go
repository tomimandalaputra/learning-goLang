package main

import "fmt"

func main() {
	var numbers [2]int
	fmt.Printf("%+v\n", numbers)

	numbers[0] = 1
	numbers[1] = 2
	fmt.Printf("%+v\n", numbers)

	primes := [4]int{2, 3, 5, 7}
	fmt.Printf("%+v\n", primes)

	primes[3] = 11
	fmt.Printf("%+v\n", primes)

	for i := range len(primes) {
		fmt.Printf("%d\n", primes[i])
	}

	// array 2 dimensi
	var matrix [2][3]int
	fmt.Printf("%+v\n", matrix)

	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 5
	matrix[1][1] = 3
	fmt.Printf("%+v\n", matrix)
}
