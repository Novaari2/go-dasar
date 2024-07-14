package helper

var version = "1.0.0" // tidak bisa diakses dari luar package karena awalannya menggunakan huruf kecil
var Application = "golang"

//tidak bisa diakses dari luar package
func sayGoodBye(name string) string {
	return "Good Bye" + name
}

// bisa diakses dari luar package
func SayHello(name string) string {
	return "Hello " + name
}
