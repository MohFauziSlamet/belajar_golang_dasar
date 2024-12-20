package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World")
}

func main() {
	/*
		Function
		Sebelumnya kita sudah mengenal sebuah function yang wajib
		dibuat agar program kita bisa berjalan, yaitu function main.
		Function adalah sebuah blok kode yang sengaja dibuat dalam
		program agar bisa digunakan berulang-ulang.
		Cara membuat function sangat sederhana,
		hanya dengan menggunakan kata kunci func,
		lalu diikuti dengan nama function-nya
		dan blok kode isi function nya.
		Setelah membuat function, kita bisa mengeksekusi
		function tersebut dengan memanggilnya menggunakan
		kata kunci nama function nya diikuti tanda kurung buka, kurung tutup.

	*/

	/// Memanggil func secara langsung
	sayHello()
	sayHello()
	sayHello()

	println()

	/// Memanggil func dengan perulangan
	for i := 0; i <= 5; i++ {
		fmt.Print(i+1, " ")
		sayHello()
	}

}
