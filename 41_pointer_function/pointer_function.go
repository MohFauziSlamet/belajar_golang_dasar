package main

import (
	"fmt"
)

// struct
type Address struct {
	city, province, country string
}

// funtion biasa (bukan parameter pointer )(1)
func ChangeAddressToIndonesia(address Address) {
	address.country = "Indonesia"
}

// funtion  (parameter pointer )(2)
func ChangeAddressToIndonesiaPointerParams(address *Address) {
	address.country = "Indonesia"
}

func main() {
	fmt.Println()

	/*
		Pointer di Function.
		● Saat kita membuat parameter di function, secara default
		adalah pass by value, artinya data akan dicopy lalu dikirim ke
		function tersebut.
		● Oleh karena itu, jika kita mengubah data di dalam function,
		data yang aslinya tidak akan pernah berubah.
		● Hal ini membuat variable menjadi aman, karena tidak akan
		bisa diubah.
		● Namun kadang kita ingin membuat function yang bisa mengubah
		data asli parameter tersebut.
		● Untuk melakukan ini, kita juga bisa menggunakan pointer di
		function.
		● Untuk menjadikan sebuah parameter sebagai pointer,
		kita bisa menggunakan operator * diparameternya.
	*/

	// contoh func biasa (1)
	address1 := Address{city: "Malang", province: "Jawa Timur"}

	// menjalankan change to indonesia
	ChangeAddressToIndonesia(address1)

	/*
		Apakah akan bisa merubah data variable asli.
	*/
	fmt.Println(address1)
	/*
		Hasilnya, variable asli tidak berubah, dan menghasilkan
		output {Malang Jawa Timur }
	*/

	fmt.Println("======================================================================")
	fmt.Println()

	address2 := Address{city: "Batu", province: "Jawa Timur"}

	// menjalankan change to indonesia with pointer (2)
	ChangeAddressToIndonesiaPointerParams(&address2)

	/*
		Apakah akan bisa merubah data variable asli.
	*/
	fmt.Println(address2)
	/*
		Hasilnya, variable asli berubah, dan menghasilkan
		output {Batu Jawa Timur Indonesia}.


		parameter nya juga bisa dengan di deklarasikan terlebih dahulu
		seperti berikut.
	*/

	address3 := Address{city: "Suhat", province: "Jawa Timur"}

	// var address3Pointer *Address = &address3 // sama
	address3Pointer := &address3 // sama

	ChangeAddressToIndonesiaPointerParams(address3Pointer)

	fmt.Println(address3)

	/*
		Hasil outputnya :{Suhat Jawa Timur Indonesia}.
		artinya , didalam funtion kita bisa mengubah nilai dari
		variable asli.

	*/

}
