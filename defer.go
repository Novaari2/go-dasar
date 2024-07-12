package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("Run Application")
}

func main() {
	/**
		mudahnya hampir sama seperti try catch finally
		Defer: defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai dieksekusi
		Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi
	**/

	runApplication()
}
