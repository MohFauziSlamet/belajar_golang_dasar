package main

import "fmt"

func main() {
	/*
		Variable
		Variable adalah tempat untuk menyimpan data.
		Variable digunakan agar kita bisa mengakses
		data yang sama dimanapun kita mau.
		Di Go-Lang Variable hanya bisa menyimpan tipe data yang sama,
		jika kita ingin menyimpan data yang berbeda-beda jenis,
		kita harus membuat beberapa variable.
		Untuk membuat variable, kita bisa menggunakan kata kunci var,
		lalu diikuti dengan nama variable dan tipe datanya.

	*/

	var name string
	name = "Moh Fauzi"
	fmt.Println(name)

	name = "Moh Fauzi Slamet"
	fmt.Println(name)

	// ==============================================================================
	fmt.Println("==============================================================================")

	/*
		Tipe Data Variable
		Saat kita membuat variable, maka kita wajib
		menyebutkan tipe data variable tersebut.
		Namun jika kita langsung menginisialisasikan data pada variable nya,
		maka kita tidak wajib menyebutkan tipe data variable nya.
	*/
	var friendName = "Moh Fauzi"
	fmt.Println(friendName)

	friendName = "Moh Fauzi Slamet"
	fmt.Println(friendName)

	var age = 999
	fmt.Println(age)

	// ==============================================================================
	fmt.Println("==============================================================================")

	/*
		Kata Kunci Var
		Di Go-Lang, kata kunci var saat membuat variable tidak lah wajib.
		Asalkan saat membuat variable kita langsung menginisialisasi datanya.
		Agar tidak perlu menggunakan kata kunci var, kita perlu menggunakan
		kata kunci := saat menginisialisasikan data pada variable tersebut.
	*/
	girlFrined := "Moh Fauzi"
	fmt.Println(girlFrined)

	girlFrined = "Moh Fauzi Slamet"
	fmt.Println(girlFrined)

	long := 999
	fmt.Println(long)
	// ==============================================================================
	fmt.Println("==============================================================================")

	/*
		Deklarasi Multiple Variable
		Di Go-Lang kita bisa membuat variable secara sekaligus banyak
		Code yang dibuat akan lebih bagus dan mudah dibaca
	*/

	var (
		firstName string = "Moh Fauzi" // dengan menyebut secara spesifik type nya.
		lastName         = "Slamet"    // tanpa menyebut type nya. semua boleh.
	)
	fmt.Println(firstName, lastName)

}
