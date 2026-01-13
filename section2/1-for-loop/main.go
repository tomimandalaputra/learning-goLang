package main

import "fmt"

func main() {
	// for -- only way to loop
	fmt.Println("========= C-style Loop =========")
	for i := 1; i <= 4; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n========= while style =========")
	k := 3
	for k > 0 {
		fmt.Println(k)
		k--
	}

	fmt.Println("\n========= infinite loop =========")
	counter := 0
	for {
		fmt.Println("Counter:", counter)
		counter++

		if counter >= 4 {
			break
		}
	}

	fmt.Println("\n========= skipping =========")
	for j := 1; j <= 10; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
	}

	fmt.Println("\n========= array =========")
	foods := [3]string{"Bakso", "Mie ayam", "Gudeg"}

	// for index, value := range foods {
	// 	fmt.Println(index, value)
	// }

	// for _, value := range foods {
	// 	fmt.Println(value)
	// }

	for index := range foods {
		fmt.Println(foods[index])
	}
}
