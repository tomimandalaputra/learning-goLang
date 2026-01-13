package main

import "fmt"

/**
* iota merupakan mesin penghitung otomatis digunakan dalam block const
* dimulai dari angka 0
 */

// define enum bernasarkan generate value otomatis by mechine
const (
	Sunday = iota
	Monday
	Thuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// define enum secara ekplisit jelas nilainya
// const (
// 	Sunday    = 0
// 	Monday    = 1
// 	Thuesday  = 2
// 	Wednesday = 3
// 	Thursday  = 4
// 	Friday    = 5
// 	Saturday  = 6
// )

func main() {
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Thuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
}
