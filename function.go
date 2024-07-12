package main

import "fmt"

// function return multiple value
func getFullname() (string, string) {
	return "nova", "ariyanto"
}

// named return values
func getClubFav() (name, address, country string, standing int) {
	name = "chelsea"
	address = "london"
	country = "England"
	standing = 1

	return name, address, country, standing
}

// variadic function
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

// function as value
func goodBye(name string) string {
	return "Good Bye " + name
}

// function as parameter
func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello ", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

// function type declaration
type Filter func(string) string

func helloFilter(name string, filter Filter) {
	fmt.Println("Hello ", filter(name))
}

func main() {
	firtsname, lastname := getFullname()
	_, namaBelakang := getFullname()

	fmt.Println(firtsname)
	fmt.Println(lastname)

	fmt.Println(namaBelakang)

	club, address, country, standing := getClubFav()

	fmt.Println(club)
	fmt.Println(address)
	fmt.Println(country)
	fmt.Println(standing)

	total := sumAll(2, 3, 4, 5, 6)
	fmt.Println(total)

	// slice sebagai parameter
	numbers := []int{10, 10, 10, 10, 10}
	total = sumAll(numbers...)
	fmt.Println(total)

	//implement func as value
	googbye := goodBye
	fmt.Println(googbye("nova"))

	// implement func as parameter
	sayHelloWithFilter("nova", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)

	helloFilter("Anjing", spamFilter)
}
