package main

import "fmt"

func Sum[T int | float32 | float64](numbers ...T) T {
	var total T
	for _, number := range numbers {
		total += number
	}
	return total
}

type Number interface {
	float32 | float64
}

func CalculatePPN[T Number](subTotal T) T {
	return subTotal * 0.1
}

func main() {
	dataInt := []int{1, 2, 3, 4, 5}
	dataFloat := []float64{1.1, 1.4, 1.5}
	fmt.Println(Sum(dataInt...))
	fmt.Println(Sum(dataFloat...))

	dataRandom := Sum(1.0, 2.25, 3.5)
	fmt.Println(dataRandom)

	fmt.Println("\n---- PPN ----")
	subTotal := Sum(dataFloat...)
	resultPPN := CalculatePPN(subTotal)
	total := Sum(subTotal, resultPPN)

	fmt.Println(total)
}
