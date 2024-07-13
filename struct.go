package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	/**
	 * struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
	 * struct biasanya reperesentasi data dalam program aplikasi yang kita buat
	 * data di struct disimpan dalam field
	 * sederhananya, struct adalah kumpulan dari field
	 * struct merupakan prototype data
	 * struct tidak bisa langsung digunakan
	 * namun kita bisa membuat data/object dari struct yg telah kita buat
	 */

	var nova Customer
	nova.Name = "Nova Ariyanto"
	nova.Address = "Semarang"
	nova.Age = 26

	fmt.Println(nova)

	// struct literal
	club := Customer{
		Name:    "chelsea",
		Address: "London",
		Age:     100,
	}

	fmt.Println(club)

	// atau

	john := Customer{"john lennon", "london", 40} // urutannya harus sesuai

	fmt.Println(john)
}
