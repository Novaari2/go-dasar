package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	nova := Man{"Nova"}
	nova.Married()

	fmt.Println(nova.Name)
}
