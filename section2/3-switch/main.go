package main

import "fmt"

func main() {
	today := "Sunday"

	switch today {
	case "Saturday", "Sunday":
		fmt.Println("Weekend!, no work")
	case "Monday", "Thuesday":
		fmt.Println("Work days, Lots of meetings")
	default:
		fmt.Println("Mid-week")
	}

	/**
	* interface{} berarti “bisa berisi nilai dg tipe apa pun” -> (any)
	* sedangkan i.(type) di dalam switch adalah type switch
	 */
	checkType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("Integer: %d\n", v)
		case string:
			fmt.Printf("String: %s\n", v)
		case bool:
			fmt.Printf("Boolean: %t\n", v)
		default:
			fmt.Printf("Unknown type: %T\n", v)
		}
	}

	/**
	* format specifier milik fmt.Printf (%d, %s, %t, %T)
	* %d  → decimal integer
	* %s  → string
	* %t  → boolean
	* %T  → type dari nilai
	* specifier bisa membantu kita dalam debugging sesuai kebutuhan
	 */

	checkType(2)       // Integer: 2
	checkType("Hello") // String: Hello
	checkType(true)    // Boolean: true
	checkType(3.14)    // Unknown type: float64
}
