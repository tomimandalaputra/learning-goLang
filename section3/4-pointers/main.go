package main

import "fmt"

func modifyValue(val int) {
	val *= 10
	fmt.Printf("modifyValue: %+v\n", val)
}

func modifyPointer(val *int) {
	if val == nil {
		fmt.Println("Val is nil")
		return
	}

	*val = *val * 10 // dereferencing
	fmt.Printf("modifyPointer: %+v\n", val)
}

func main() {
	num := 10
	modifyValue(num)
	fmt.Println(num)

	modifyPointer(&num)
	fmt.Println(num)

	grade := 80
	gradePtr := &grade
	fmt.Printf("gradePtr: %+v\n", gradePtr)
	fmt.Printf("new Pointer: %+v\n", &gradePtr)
	fmt.Printf("dereference: %+v\n", *(&gradePtr))
}
