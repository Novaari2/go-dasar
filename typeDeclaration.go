package main

import "fmt"

func main() {
	// type declarattion adalah kemampuan membuat ulant tipe data baru dari tipe data yang sudah ada
	// biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada

	type noKTP string

	var ktpNova noKTP = "33074623778238"
	fmt.Println(ktpNova)
	fmt.Println(noKTP("3472937472927"))
}
