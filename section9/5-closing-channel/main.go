package main

import "fmt"

/*
Diasumsikan worker tidak tahu berpa jumlah job yang akan dijalankan
Func worker hanya melakukan:
1. Mengabil 1 nilai dari jobs
2. Lalu memprosesnya
4. Selanjutnya akan mengirikan ke dalam channel result
5. Ulangi kembali sampai channel jobs ditutup
*/
func worker(id int, jobs chan int, results chan int) {
	for job := range jobs {
		fmt.Println("worker", id, "processing job", job)
		results <- job
	}
}

func main() {

	// define channel jobs dan channel results
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	/*
		Pd kode berikut ini kita bisa membayangkan
		ada 5 (channel jobs) tugas tetapi hanya ada 3 pegawai (maks iterasi pd worker)
		yang bekerja secara paralel.

		-> Artinya akan ada 3 job dikerjakan secara bersamaan,
		sisahnya menunggu di channel jobs
	*/
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	/*
		Selanjutnya, melakukan pengiriman 5 job (maks jml iterasi)
		stau per satu ke channel jobs sampai channel jobs ditutup.

		Ingat bahwa kita hanya mempunyai 3 pegawai (worker), maka pembagiannya akan dikerjakan
		tanpa urutan yang pasti/ karena scheduler go yang akan menentukan
	*/
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	/*
		Karena ada 5 job yang dikirimkan, maka akan ada 5 hasil.
		Hal ini bergantung pada jumlah job bukan jumlah worker!
	*/
	for k := 1; k <= 5; k++ {
		fmt.Printf(" - Result: %d\n", <-results)
	}

	/*
		CONTOH OUTPUT 1
		worker 3 processing job 2
		worker 3 processing job 4
		worker 3 processing job 5
			- Result: 2
			- Result: 4
			- Result: 5
		worker 1 processing job 1
			- Result: 1
		worker 2 processing job 3
			- Result: 3

		CONTOH OUTPUT 2
		worker 2 processing job 3
		worker 2 processing job 4
		worker 2 processing job 5
		 - Result: 3
		 - Result: 4
		 - Result: 5
		worker 1 processing job 2
		 - Result: 2
		worker 3 processing job 1
		 - Result: 1
	*/
}
