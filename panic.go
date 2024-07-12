package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("ERROR")
	}
}

func main() {
	/**
	 * Panic adalah function yang dapat kita gunakan untuk menghentikan program
		panic biasanya dipanggil ketika terjadi panic saat program kita berjalan
		saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi
	*/

	runApp(true)
}
