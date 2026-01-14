package main

import "fmt"

func greeting(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func sum(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}

func calculateArea(w, h float64) float64 {
	if w < 0 || h < 0 {
		fmt.Println("Error: width or hight must be positive")
		return 0.0
	}

	return w * h
}

func main() {
	greeting("Tomi")
	sum(1, 4)

	area := calculateArea(3.14, 2)
	fmt.Println(area)
}
