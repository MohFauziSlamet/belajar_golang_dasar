package main

import "fmt"

func main() {
	/*
		Constant
		Constant adalah variable yang nilainya tidak bisa diubah lagi
		setelah pertama kali diberi nilai.
		Cara pembuatan constant mirip dengan variable,
		yang membedakan hanya kata kunci yang digunakan adalah const, bukan var.
		Saat pembuatan constant, kita wajib langsung menginisialisasikan datanya.
	*/

	const name string = "Moh Fauzi"
	fmt.Println(name)

	// Akan error ketika nilai nya diubah
	// name = "Moh Fauzi Slamet"

	// ==============================================================================
	fmt.Println("==============================================================================")

	/*
		Deklarasi Multiple Constant
		Sama seperti variable, di Go-Lang  juga kita bisa membuat constant
		secara sekaligus banyak.
	*/
	const (
		firstName = "Moh Fauzi"
		lastName  = "Slamet"
		value     = 1000
	)

	/*
		Ketika kita mendeklarasikan constan , tidak akan terjadi error , meskipun tidak digunakan.
	*/
	fmt.Println(firstName, lastName, value)

}
