package main

import "fmt"

/*
Go memiliki interface bawaan yaitu Stringer yang digunakan untuk
mengubah struktur data menjadi string agar mudah dibaca.
*/

type Product struct {
	Name  string
	Price int
}

func (p Product) String() string {
	return fmt.Sprintf("Product: %s, Price: %d", p.Name, p.Price)
}

func main() {
	product := Product{Name: "Laptop", Price: 1000000}
	fmt.Println(product)

	// --- output ---
	// sebelum implementasi Stringer: {Laptop 1000000}
	// setelah implementasi Stringer: Product: Laptop, Price: 1000000
}
