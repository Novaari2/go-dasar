package main

import "fmt"

func main() {
	var name string            // wajib menyebutkan tipe datanya
	var umur = 26              // tidak perlu disebutkan tipe datanya karena langsung di inisiasi (otomatis dideteksi)
	alamat := "jalan pisang 1" // jika menggunakan := tidak perlu menggunakan kata kunci var

	//multiple variable
	var (
		firstName = "Nova"
		lastName  = "Ariyanto"
	)
	//end

	name = "nova"

	fmt.Println(name)
	fmt.Println(umur)
	fmt.Println(alamat)
	fmt.Println(firstName)
	fmt.Println(lastName)
}

// di golang, variable hanya bisa menyimpan tipe data yang sama
