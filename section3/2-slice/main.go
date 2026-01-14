package main

import "fmt"

func main() {
	names := []string{"Tomi", "Mandala", "Putra"}
	fmt.Println(names)

	// make([]int, length, capacity)
	items := make([]int, 3, 5)
	fmt.Printf("Items: %+v, Length: %d, Capacity: %d\n", items, len(items), cap(items))

	items = append(items, 1)
	items = append(items, 2)
	items = append(items, 3)
	items = append(items, 4)

	// [0,0,0,1,2,3,4]
	fmt.Printf("Items: %+v, Length: %d, Capacity: %d\n", items, len(items), cap(items))

	// Slice expression Mulai dari index 3 sampai sebelum index 5
	fmt.Printf("result: %+v", items[3:5]) // output [1 2]
}
