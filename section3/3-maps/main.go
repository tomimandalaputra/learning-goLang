package main

import "fmt"

func main() {
	studentGrades := map[string]int{
		"Tomi":    90,
		"Mandala": 80,
		"Putra":   75,
	}
	fmt.Println(studentGrades)

	studentGrades["Tomi"] = 98
	fmt.Println(studentGrades)

	key := "Putra"
	if value, ok := studentGrades[key]; ok {
		fmt.Printf("%s: %d, status ok (when key is match): %t\n", key, value, ok)
	}

	delete(studentGrades, "Putra")
	fmt.Println(studentGrades)

	// var configs map[string]int // -> nill
	configs := make(map[string]int) // -> empty but valid
	fmt.Printf("Result: %+v, Type: %T\n", configs, configs)

	if configs == nil {
		fmt.Println("Config is nil")
	}
}
