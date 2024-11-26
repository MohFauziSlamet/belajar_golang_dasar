package main

import "fmt"

func getFullName() (firstName string, lastName string) {
	firstName = "Moh Fauzi "
	lastName = "Slamet"
	return

}

func main() {
	/*
		Named Return Values
		Biasanya saat kita memberi tahu bahwa 
		sebuah function mengembalikan value,
		maka kita hanya mendeklarasikan tipe 
		data return value di function.
		Namun kita juga bisa membuat variable secara 
		langsung di tipe data return function nya.

	*/

	firstName, lastName := getFullName()

	fmt.Println(firstName, lastName)

	fmt.Println("==============================================================================")

	/*
		Menghiraukan Return Value.
		Multiple return value wajib ditangkap semua value nya.
		Jika kita ingin menghiraukan return value tersebut, 
		kita bisa menggunakan tanda _ (garis bawah).

	*/

	firstName2, _ := getFullName()

	fmt.Println(firstName2)

}
