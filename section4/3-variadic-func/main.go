package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

/**
* variadic parameter
* fungsi menerima nol atau lebih int
* di dalam fungsi sumTotal, numbs bertipe []int (slice)
 */
func sumTotal(numbs ...int) int {
	total := 0

	for _, num := range numbs {
		total += num
	}

	return total
}

func main() {
	result := sum(1, 2)
	fmt.Println(result)

	// option 1
	// totalValue := sumTotal(1, 2, 3, 4, 5)

	// option 2
	numbers := []int{1, 2, 3, 4, 5}
	totalValue := sumTotal(numbers...)
	fmt.Println(totalValue)
}
