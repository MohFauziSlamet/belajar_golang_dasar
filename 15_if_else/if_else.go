package main

import (
	"fmt"
)

func main() {
	/*
		If Expression
		If adalah salah satu kata kunci yang digunakan untuk percabangan.
		Percabangan artinya kita bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi.
		Hampir di semua bahasa pemrograman mendukung if expression.
	*/

	/// Membuat if expresion
	name1 := "Fauzi"

	if name1 == "Fauzi" {
		fmt.Println("Helo ", name1)
	}

	fmt.Println("======================================================================")

	/*
		Else Expression
		Blok if akan dieksekusi ketika kondisi if bernilai true.
		Kadang kita ingin melakukan eksekusi program tertentu jika kondisi if bernilai false.
		Hal ini bisa dilakukan menggunakan else expression.
	*/

	/// Membuat if else expresion
	name2 := "Fauzii"

	if name2 == "Fauzi" {
		fmt.Println("Helo ", name2)
	} else {
		fmt.Println("Apa boleh kenalan ")
	}

	fmt.Println("======================================================================")

	/*
		Else If Expression
		Kadang dalam If, kita butuh membuat beberapa kondisi.
		Kasus seperti ini, kita bisa menggunakan Else If expression.
	*/

	/// Membuat if else expresion
	name3 := "Ludy"

	if name3 == "Fauzi" {
		fmt.Println("Helo ", name3)
	} else if name3 == "Ludy" {
		fmt.Println("Helo ", name3)

	} else {
		fmt.Println("Apa boleh kenalan ")
	}

	fmt.Println("======================================================================")

	/*
		If dengan Short Statement
		If mendukung short statement sebelum kondisi.
		Hal ini sangat cocok untuk membuat statement yang sederhana sebelum melakukan pengecekan terhadap kondisi.

	*/

	if lenght := len(name1); lenght > 4 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
