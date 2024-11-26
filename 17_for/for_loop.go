package main

import (
	"fmt"
)

func main() {
	/*
		For Expression
		Dalam bahasa pemrograman, biasanya ada fitur 
		yang bernama perulangan.
		Salah satu fitur perulangan adalah for loops.
	*/

	/// Membuat for expresion
	counter := 1

	for counter <= 5 {
		fmt.Println("Index ke", counter)
		counter++
	}

	fmt.Println("======================================================================")
	fmt.Println()

	/*
		For dengan Statement
		Dalam for, kita bisa menambahkan statement,
		dimana terdapat 2 statement yang bisa tambahkan di for.
		Init statement, yaitu statement sebelum for di eksekusi.
		Post statement, yaitu statement yang akan selalu 
		dieksekusi di akhir tiap perulangan.
	*/

	/// i :=0 => merupakan init statement.
	/// i ++ => merupakan post statement.
	for i := 0; i <= 5; i++ {
		fmt.Println("Index ke", i)
	}
	fmt.Println("======================================================================")
	fmt.Println()

	/*
		For juga bisa digunakan untuk mengakses 
		data pada array , slice dan map.
	*/

	/// Array
	fmt.Println("\nLoop in Array")
	array := [...]string{"Moh", "Fauzi", "Slamet"}
	fmt.Println("Panjang array", len(array))
	fmt.Println("Panjang array", cap(array))

	for i := 0; i < len(array); i++ {
		fmt.Println("Index ke", i, "=", array[i])
	}

	fmt.Println("======================================================================")
	fmt.Println()

	/// Slice
	fmt.Println("Loop in Slice")
	// slice := []string{"Moh", "Fauzi", "Slamet"}
	slice := []string{"Moh", "Fauzi", "Slamet"}
	fmt.Println("Panjang slice", len(slice))
	fmt.Println("Capasity slice", cap(slice))

	for i := 0; i < len(slice); i++ {
		fmt.Println("Index ke", i, "=", slice[i])
	}

	fmt.Println("======================================================================")
	fmt.Println()

	/*
		Range
		peruangan dengan menggunakan range.
	*/

	for index, value := range array {
		if index == 0 {
			fmt.Println("Perulangan Array dengan Range")
		}
		fmt.Println("Index ke", index, "=", value)
	}

	fmt.Println()

	for index, value := range slice {
		if index == 0 {
			fmt.Println("Perulangan Slice dengan Range")
		}
		fmt.Println("Index ke", index, "=", value)
	}

	fmt.Println()

	person := map[string]string{
		"name":    "fauzi",
		"address": "Malang",
	}
	fmt.Println("Perulangan Map dengan Range")
	for key, value := range person {
		fmt.Println("key", key, ":", value)
	}

	fmt.Println("======================================================================")
	fmt.Println()
	/*
		didalam for range terdapat index atau key.
		jika kita tidak memakai data tersebut,
		bisa kita hilangkan dengan mengganti dengan 
		tanda underscore ( _ ).
	*/
	fmt.Println("Perulangan Array dengan Range")
	for _, value := range array {
		fmt.Println("Tanpa Index", "=", value)
	}

	fmt.Println()

	fmt.Println("Perulangan Slice dengan Range")
	for _, value := range slice {

		fmt.Println("Tanpa index", "=", value)
	}

	fmt.Println()

	fmt.Println("Perulangan Map dengan Range")
	for _, value := range person {
		fmt.Println("Tanpa key", ":", value)
	}

}
