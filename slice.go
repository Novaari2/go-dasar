package main

import "fmt"

func main() {
	// slice adalah potongan dari data array
	// ukurannya bisa berubah
	// slice dan array selalu terkoneksi, slice adalah data yang mengakses sebagian atau seluruh data di array
	/**
		Tipe data slice memiliki 3 data, yaitu pointer, length, capacity
		Pointer adalah penunjuk data pertama di array pada slice
		Length adalah panjang dari slice
		Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

		## function yang ada di slice
		len(slice) -> untuk mendapatkan panjang
		cap(slice) -> untuk mendapatkan kapasitas
		append(slice, data) -> membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh maka akan membuat array baru
		make([]TypeData, length, capacity) -> membuat slice baru
		copy(destination, source) -> menyalin slice dari source ke destination

	**/
	names := [...]string{"chelsea", "bayern", "barca", "madrid", "inter", "milan"}
	slice1 := names[4:6]

	fmt.Println(slice1[0])
	fmt.Println(slice1[1])

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	// append slice

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jum'at", "sabtu", "minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu baru"
	daySlice1[1] = "Minggu baru"
	// saat merubah isi slice, akan merubah data di arraynya

	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur baru")
	daySlice2[0] = "Upsss"

	fmt.Println(daySlice2)
	fmt.Println(days)

	// membuat slice baru
	newSlice := make([]string, 2, 5)
	newSlice[0] = "nova"
	newSlice[1] = "ari"
	//newSlice[2] = "yanto" // error, harusnya menggunakan appeng

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "yanto")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	// copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// perbedaan array dan slice
	iniArray := [...]int{1, 2, 3, 4, 5} // tentukan jumlah arraynya atau gunakan ...
	iniSlice := []int{1, 2, 3, 4, 5}    // tidak perlu menentukan jumlahnya

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
