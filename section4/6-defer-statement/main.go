package main

import "fmt"

func simpleDefer() {
	/*
		Secara normal kode dijalankan dari atas ke bawah, tetapi
		pada defer memiliki aturan khusus yakni:
		-> baris yang diawali defer tidak langsung dieksekusi
		akan disimpan didalam stack defer, akan dijalankan saat
		fungsi akan selesai.

		lalu bagaimana dengan defer lebih dari satu dlm sebuah function?
	*/
	fmt.Println("Func simpleDefer: Start")
	defer fmt.Println("Func simpleDefer: deferred")
	fmt.Println("Func simpleDefer: Middle")
	fmt.Println("Func simpleDefer: Middle")
	fmt.Println("Func simpleDefer: Middle")
}

func lifoSimpleDefer() {
	/*
		LIFO (Last In First Out)
		--- analogi ---
		Bayangkan kamu punya dua catatan (First) dan (Second) maka, defer akan dijalankan
		berdasarkan urutan penumpukan tesebut ke dalam stack defer.
		Gini, ada contoh dua baris defer dlm kode. Pertama Go bertemu defer First, maka
		catatan First masuk ke tumpukan. Lalu bertemu defer Second, catatan Second ditaruh di atas First.
		Saat fungsi mau keluar, Go mulai membersihkan dari atas tumpukan,
		maka Second dieksekusi dulu, baru First.
	*/
	fmt.Println("Func lifoSimpleDefer: Start")
	defer fmt.Println("First: deferred")
	defer fmt.Println("Second: deferred")
	fmt.Println("Func simpleDefer: Middle")
}

func main() {
	// simpleDefer()
	lifoSimpleDefer()

	// --- OUTPUT ---
	// Func lifoSimpleDefer: Start
	// Func simpleDefer: Middle
	// Second: deferred
	// First: deferred
}
