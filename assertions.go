package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	/**
	 * type assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
	 * fitur ini sering digunakan ketika kita bertemu dengan data interface kosong
	 */

	var result = random()
	var resultString = result.(string)
	fmt.Println(resultString)

	// jika tipe data kembalian tidak sesuai maka akan terjadi panic
	//var resultInt = result.(int) // panic karena tipe data tidak sesuai
	//fmt.Println(resultInt)

	// best practice
	switch value := result.(type) { //lakukan pengecekan tipe datanya dulu
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}
}
