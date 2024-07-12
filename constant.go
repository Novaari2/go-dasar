package main

import "fmt"

func main() {
	const usia = 26

	//multiple constant
	const (
		first = 1
		last  = 10
	)

	fmt.Println(usia)
	fmt.Println(first)
	fmt.Println(last)
}
