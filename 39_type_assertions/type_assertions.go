package main

import (
	"fmt"
)

/*
Type Assertions
● Type Assertions merupakan kemampuan merubah tipe data menjadi
tipe data yang diinginkan.
● Fitur ini sering sekali digunakan ketika kita bertemu dengan
data interface kosong.
*/

// membuat interface
func random() interface{} {
	return "Barochatul Mauludy"
}

func main() {
	fmt.Println()

	// contoh benar
	// result := random()
	// resultStr := result.(string)

	// fmt.Println(resultStr)
	/*
		hasil nya tidak akan error, karena kita yakin dan tau bahwa
		type nya adalah string.
	*/

	// contoh salah
	// result := random()
	// resultStr := result.(int)

	// fmt.Println(resultStr)
	/*
		hasil nya akan error.
		dikarenakan type return dari random adl string.
		tetapi di convert ke dalam int.
		maka akan menampilkan pesan error sbb :
		panic: interface conversion: interface {} is string, not int.
	*/

	// cara mencegah terjadi error seperti di atas
	/*
		Type Assertions Menggunakan Switch
		● Saat salah menggunakan type assertions, maka bisa berakibat
		terjadi panic di aplikasi kita.
		● Jika panic dan tidak ter-recover, maka otomatis program kita akan mati.
		● Agar lebih aman, sebaiknya kita menggunakan switch expression
		untuk melakukan type assertions.
	*/

	result := random()

	switch value := result.(type) {
	case string:
		fmt.Println("value", value, "is string")
	case int:
		fmt.Println("value", value, "is string")
	default:
		fmt.Println("Unknown type")

	}

	fmt.Println("======================================================================")
	fmt.Println()

}
