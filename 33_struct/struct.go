package main

import "fmt"

/*
Struct
● Struct adalah sebuah template data yang digunakan untuk
menggabungkan satu atau lebih tipe data lainnya dalam
satu kesatuan.
● Struct biasanya representasi data dalam program
aplikasi yang kita buat.
● Data di struct disimpan dalam field.
● Sederhananya struct adalah kumpulan dari field.
● Dev golang, biasanya membuat nama struct dengan Upercase
contoh : DataCustomer.
● Dan juga Dev golang, biasanya membuat variable/field
didalam struct dengan Upercase
contoh : FirstName / LastName.
*/

/*
Membuat Data Struct
● Struct adalah template data atau prototype data.
● Struct tidak bisa langsung digunakan.
● Namun kita bisa membuat data/object (exp:variabel) ,
dari struct yang telah kita buat.
*/

// Membuat struct
type Customer struct {
	Name, Addres string // penulisan singkat

	// Penulisan panjang
	/*
		Name string
		Addres string
	*/

	Age int
}

func main() {
	fmt.Println()

	fmt.Println("Membuat struct 1")

	var ludy Customer

	// Memasukan data
	ludy.Name = "Barochatul Mauludy"
	ludy.Addres = "Suko"
	ludy.Age = 24

	// jika data tidak di isi, maka akan tampil kosong .
	// hal ini berlaku dengan type data yang lain.
	// jika semua data tidak di isi , maka akan tampil : {  0}
	// fmt.Println(ludy.Addres)

	fmt.Println(ludy)

	fmt.Println("======================================================================")
	fmt.Println()

	/*
		Struct Literals
		● Sebelumnya kita telah membuat data dari struct,
		namun sebenarnya ada banyak cara yang bisa kita
		gunakan untuk membuat data dari struct.

	*/

	fmt.Println("Membuat struct 2")

	// dengan memasukan data secara langsung
	mauludy := Customer{"Barochatul Mauludy", "Suko", 24}

	fmt.Println(mauludy)

	/*
		penggunaan cara ini terlalu berbahaya.
		alasan nya, jika data struct dirubah posisi deklarasinya,
		maka akan terjasi error.

		saat ini :
		Name, Addres string
		Age int

		jika dirubah seperti ini :
		Name string
		Age int
		Addres string

		maka akan terjadi error.
	*/

	fmt.Println("======================================================================")
	fmt.Println()

	fmt.Println("Membuat struct 3")

	// dengan memasukan nama variable/field
	maul := Customer{
		Name:   "Barochatul Mauludy",
		Addres: "Suko",
		Age:    24,
	}

	fmt.Println(maul)

	/*
		penggunaan cara ini lebih direkomendasi kan.
		alasan nya, mudah dibaca. misal field ini dengan value ini.
		dan apabila ada perubaha pada data struct, maka tidak akan
		terjadi error.

		saat ini :
		Name, Addres string
		Age int

		jika dirubah seperti ini :
		Name string
		Age int
		Addres string

		tidak akan terjadi error.
	*/

	fmt.Println("======================================================================")
	fmt.Println()

}
