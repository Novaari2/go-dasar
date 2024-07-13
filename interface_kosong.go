package main

import "fmt"

func ups() interface{} {
	//return 1
	//return true
	return "Ups"
}

// atau
func wow() any {
	return true
}

func main() {
	/**
	 * interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya
	 * interface kosong juga memiliki type alias bernama any
	 */

	kosong := ups()
	fmt.Println(kosong)

	kaget := wow()
	fmt.Println(kaget)
}
