package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	/**
	 * go memiliki function new yang bisa digunakan untuk membuat operator
	 * namun new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
	 */

	alamat1 := new(Address)
	alamat2 := alamat1

	alamat2.Country = "Nepal"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
