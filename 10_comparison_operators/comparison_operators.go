package main

import "fmt"

func main() {
	/*
		Operasi Perbandingan
		Operasi perbandingan adalah operasi untuk membandingkan dua buah data
		Operasi perbandingan adalah operasi yang menghasilkan nilai boolean (benar atau salah)
		Jika hasil operasinya adalah benar, maka nilainya adalah true
		Jika hasil operasinya adalah salah, maka nilainya adalah false

		Operator Perbandingan
		Operator          	Keterangan
		>					Lebih Dari
		<					Kurang Dari
		>=					Lebih Dari Sama Dengan
		<=					Kurang Dari Sama Dengan
		==					Sama Dengan
		!=					Tidak Sama Dengan

	*/

	var name1 = "Fauzi"
	var name2 = "fauzi"
	println("1 ", name1 == name2)                     // false
	println("2 ", name1 != name2)                     // true
	println("3 ", name1[0], name1 <= name2, name2[0]) // true
	println("4 ", name1[0], name1 >= name2, name2[0]) // false
	println("5 ", name1[0], name1 > name2, name2[0])  // false
	println("6 ", name1[0], name1 < name2, name2[0])  // true

	/*
		Jika menggunakan <= , >= , < , > pada tipe data string,
		maka akan di cek pada kode byte index[0]. lalu kode byte index[0]
		akan dibandingkan dengan kode byte index[0] juga.
	*/

	fmt.Println("==============================================================================")

}
