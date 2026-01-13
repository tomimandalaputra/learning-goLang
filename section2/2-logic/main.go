package main

import "fmt"

func main() {
	suhu := 25
	if suhu > 30 {
		fmt.Println("Suhu lebih dari 30")
	} else {
		fmt.Println("Suhu lebih kecil dari 30")
	}

	score := 80
	if score >= 90 {
		fmt.Println("Greade: A")
	} else if score >= 80 {
		fmt.Println("Greade: B")
	} else if score >= 70 {
		fmt.Println("Greade: C")
	} else if score >= 60 {
		fmt.Println("Greade: D")
	} else {
		fmt.Println("Failed")
	}

	userAccess := map[string]bool{
		"tomi":  true,
		"putra": false,
	}

	// hasAccess, ok := userAccess["tomi"]

	// if ok && hasAccess {
	// 	fmt.Println("Tomi can access the system")
	// } else {
	// 	fmt.Println("Access denied")
	// }

	/**
	* Setiap username dipetakan ke apakah dia punya akses atau tidak.
	* "tomi" ada dan bernilai true
	* "putra" ada dan bernilai false
	* Code berjalan
	* hasAccess, ok := userAccess["putra"]
	* Hasilnya hasAccess = false dan ok = true kenapa??????
	* hasAccess merujuk ke pada value dari si map array -> putra
	* sedangkan ok merujuk pada key "putra" memang ada didalam map array
	* jadi ok akan false jika key tidak ditemukan pada map array
	 */

	if hasAccess, ok := userAccess["putra"]; ok && hasAccess {
		fmt.Println("Tomi can access the system")
	} else {
		fmt.Println("Access denied")
	}

}
