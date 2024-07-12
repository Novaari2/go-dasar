package main

import "fmt"

func main() {
	// closure adalah kemampuan sebuah function berinteraksi dengan data - data yang ada di sekitarnya
	// harus berhati hati ketika menggunakan closure, perhatikan data dan perubahan datanya

	counter := 0
	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	increment()

	fmt.Println(counter)
}
