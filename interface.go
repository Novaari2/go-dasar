package main

import "fmt"

type Person struct {
	Name string
}

type HasName interface {
	GetName() string
}

func (person Person) GetName() string {
	return person.Name
}

func sayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	/**
	 * interface adalah tipe data abstract, dia tidak dapat di implementasi langsung
	 * sebuah interface berisikan definisi - definisi method
	 * biasanya interface digunakan sebagai kontrak
	 *
	 * setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
	 * sehingga kita tidak perlu mengimplementasikan interface secara manual
	 * hal ini agak berbeda dengan bahasa pemrograman lainnya yang ketika membuat interface, kita harus menyebutkan secara eksplisit akan menggunakan interface mana
	 */

	person := Person{Name: "nova"}
	sayHello(person)

	animal := Animal{"Dog"}
	sayHello(animal)
}
