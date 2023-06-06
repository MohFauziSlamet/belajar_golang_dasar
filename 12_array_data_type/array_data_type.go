package main

import "fmt"

func main() {
	/*
		Tipe Data Array
		Array adalah tipe data yang berisikan kumpulan data dengan tipe yang sama.
		Saat membuat array, kita perlu menentukan jumlah data yang bisa ditampung oleh Array tersebut.
		Daya tampung Array tidak bisa bertambah setelah Array dibuat

		Index di Array
		Index			Data
		0				Moh
		1				Fauzi
		2				Slamet
	*/

	var names [3]string // [3] merupakan batas max dari array tsb
	names[0] = "Moh"
	names[1] = "Fauzi"
	names[2] = "Slamet"
	// names[3] = "Slamet"  // jika memasukan lebih dari panjang yang telah
	// ditentukan , akan terjadi error.

	println(names[0], names[1], names[2])

	/*
		Membuat Array Langsung
		Di Go-Lang kita juga bisa membuat Array secara langsung saat deklarasi variable

	*/

	var names2 = [3]string{
		"Moh",
		"Fauzi",
		"Slamet",
	}

	fmt.Println(names2)

	/*
			Function Array
		 	Operasi						Keterangan
			len(array)					Untuk mendapatkan panjang Array
			array[index]				Mendapat data di posisi index
			array[index] = value		Mengubah data di posisi index

	*/

	var numbers [10]int
	fmt.Println(len(names))
	fmt.Println(len(names2))
	// meskipun belum ada valuenya , tapi panjangnya sudah bisa di ambil,
	// karena sudah di pastikan saat awal deklarasi
	fmt.Println(len(numbers))

}
