package main

import "fmt"

func main() {
	// di golang tipe data array harus ditentukan jumlah datanya
	// daya tampung array tidak bisa berubah setelah dibuat

	var names [3]string

	names[0] = "nova"
	names[1] = "ari"
	names[2] = "yanto"

	fmt.Println(names)

	// membuat array secara langsung
	var values = [3]int{
		90,
		91,
		92,
	}

	fmt.Println(values)

	/** function array:
	len() : mendapatkan panjang array,
	arr[index]: mendapat data di posisi index ke n,
	arr[index] = value : mengubah data pada posisi index ke n
	**/
}
