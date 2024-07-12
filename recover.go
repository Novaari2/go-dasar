package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi error", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("ERROR")
	}
}

func main() {
	/**
	 * Recover adalah function yang bisa kita gunakan untuk menangkap data panic
	 * dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan
	 */

	runApp(true)
	fmt.Println("Nova Ariyanto")
}
