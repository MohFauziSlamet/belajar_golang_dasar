package main

import "fmt"

func main() {
	/*
		Type Declarations
		Type Declarations adalah kemampuan membuat ulang tipe data baru dari
		tipe data yang sudah ada.
		Type Declarations biasanya digunakan untuk membuat alias terhadap tipe data
		yang sudah ada.
		dengan tujuan agar lebih mudah dimengerti.
	*/

	type noKtp string

	var noKtpFauzi noKtp = "123456789" // bertipe string

	println(noKtpFauzi)
	println(noKtp("2222222222"))

	fmt.Println("==============================================================================")

}
