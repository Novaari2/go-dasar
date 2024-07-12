package main

import "fmt"

func main() {
	// konversi tipe data
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32) // akan menghasilkan number overflow karena nilai melebihi kapasitas dari int16
	var name = "nova ariyanto"
	var e = name[0]
	var eString = string(e) // agar muncul huruf, bukan nilai byte maka harus dikonversi ke string (karena name bertipe string)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	fmt.Println(eString)
}
