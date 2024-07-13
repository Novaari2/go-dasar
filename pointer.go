package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	/**
	 * secara default di go, semua variable itu di passing by value. bukan by reference
	 * artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain sebenarnya yang dikirim adalah duplikasi value nya
	 *
	 * pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
	 * dengan pointer kita bisa melakukan pass by reference
	 */
	// pass by value

	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1 // di copy (pass by value, maka jika address 2 diubah isinya tidak akan mempengaruhi data di address1)

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)

	// pass by reference

	address3 := &address1 // pointer
	//var address4 *Address = &address1 // penulisan versi panjangnya

	address3.City = "Semarang"

	fmt.Println(address1)
	fmt.Println(*address3)

	// operator asterisk
	// mengubah semua yg mengacu ke address
	*address3 = Address{"Buleleng", "Bali", "Indonesia"}
	// semua yang mengacu ke address akan berubah
	fmt.Println(address1)
	fmt.Println(address3)
}
