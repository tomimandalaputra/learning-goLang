package main

import "fmt"

func factorialRecursive(n int) int {
	if n <= 1 {
		return 1
	}

	return n * factorialRecursive(n-1)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	number := 5
	result := factorialRecursive(number)
	fmt.Println("Factorial of", number, "is", result)

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
