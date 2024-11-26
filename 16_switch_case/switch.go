package main

import (
	"fmt"
)

func main() {
	/*
		Switch Expression
		Selain if expression, untuk melakukan percabangan, 
		kita juga bisa menggunakan Switch Expression.
		Switch expression sangat sederhana dibandingkan if.
		Biasanya switch expression digunakan untuk melakukan 
		pengecekan ke kondisi dalam satu variable.
	*/

	/// Membuat switch expresion
	name1 := "Ludy"

	switch name1 {
	case "Fauzi":
		fmt.Println("Hai Fauzi")
	case "Ludy":
		fmt.Println("Hai Ludy")
	default:
		fmt.Println("Hai boleh kenalan")

	}
	fmt.Println("======================================================================")

	/*
		Switch dengan Short Statement.
		Sama dengan If, Switch juga mendukung short statement sebelum variable yang akan di cek kondisinya

	*/

	switch lenght := len(name1); lenght < 5 {
	case true:
		fmt.Println("Nama sudah sesuai")
	case false:
		fmt.Println("Nama teralu panjang")

	}

	fmt.Println("======================================================================")

	/*
		Switch Tanpa Kondisi
		Kondisi di switch expression tidak wajib.
		Jika kita tidak menggunakan kondisi di switch expression,
		kita bisa menambahkan kondisi tersebut di setiap case nya.
	*/
	lenght := len(name1)
	switch {
	case lenght > 10:
		fmt.Println("Nama terlalu panjang")

	case lenght > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah sesuai")
	}

}
