package main

import "fmt"

func getFullName() (string, string) {
	return "Moh Fauzi ", "Slamet"

}

func main() {
	/*
		Returning Multiple Values
		Function tidak hanya dapat mengembalikan satu value,
		tapi juga bisa multiple value.
		Untuk memberitahu jika function mengembalikan multiple value,
		kita harus menulis semua tipe data return value nya di function.
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
