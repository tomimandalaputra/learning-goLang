package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- Advanced Slicing Operations ---")
	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Original: %v, length: %d, capacity: %d\n", original, len(original), cap(original))

	firstSlice := original[2:5] //slice[low:high] -> low inclusive, high exclusive

	// Original -> [2 3 4]
	// length -> len = 5 - 2 = 3
	// capacity -> cap = cap(original) - 2 = 8
	fmt.Printf("Original: %v, length: %d, capacity: %d\n", firstSlice, len(firstSlice), cap(firstSlice))

	secondSlice := original[:4] // -> by default is original[0:4]
	fmt.Printf("Original: %v, length: %d, capacity: %d\n", secondSlice, len(secondSlice), cap(secondSlice))

	thirdSlice := original[6:] // -> the code original[6:len(original)] -> original[6:10]
	fmt.Printf("Original: %v, length: %d, capacity: %d\n", thirdSlice, len(thirdSlice), cap(thirdSlice))

	// OUTPUT
	// Original: [2 3 4], length: 3, capacity: 8
	// Original: [0 1 2 3], length: 4, capacity: 10
	// Original: [6 7 8 9], length: 4, capacity: 4

	fourthSlice := original[:]
	fmt.Printf("Original: %v, length: %d, capacity: %d\n", fourthSlice, len(fourthSlice), cap(fourthSlice))

	// slices.Contains(fourthSlice, 8)
	fourthSlice = append(fourthSlice, 1000)
	fmt.Printf("Original: %v, length: %d, capacity: %d\n", fourthSlice, len(fourthSlice), cap(fourthSlice))
}
