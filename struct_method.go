package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, " my name is", customer.Name)
}

func main() {
	/**
	 * struct adalah tipe data seperti tipe data lainnya, dia bisa digunanakan sebagai parameter untuk function
	 */

	ari := Customer{Name: "Nova Ari"}
	ari.sayHello("budi")
}
