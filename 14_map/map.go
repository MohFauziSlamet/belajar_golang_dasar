package main

import (
	"fmt"
)

func main() {
	/*
		Tipe Data Map
		Pada Array atau Slice, untuk mengakses data, kita menggunakan index Number dimulai dari 0.
		Map adalah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan
		jenis tipe data index yang akan kita gunakan.
		Sederhananya, Map adalah tipe data kumpulan key-value (kata kunci - nilai), dimana kata
		kuncinya bersifat unik, tidak boleh sama.
		Berbeda dengan Array dan Slice, jumlah data yang kita masukkan ke dalam Map boleh sebanyak-banyaknya,
		asalkan kata kunci nya berbeda.
		jika kita gunakan kata kunci sama, maka secara otomatis data sebelumnya akan diganti dengan data baru.
	*/

	/// Membuat map
	person := map[string]string{
		"name":    "fauzi",
		"address": "Malang",
	}
	fmt.Println("Map Person : ", person)

	/// Menambah key pada map
	person["age"] = "28"
	fmt.Println("Map Person setelah ditambahkan key baru : ", person)

	/// Mengakkses data pada map berdasarkan key
	fmt.Println("Key name :", person["name"])
	fmt.Println("Key address :", person["address"])
	fmt.Println("Key age :", person["age"])

	/*
			Function Map
		 	Operasi						Keterangan
			len(map)					:Untuk mendapatkan jumlah data di map
			map[key]					:Mengambil data di map dengan key
			map[key] = value			:Mengubah data di map dengan key
			make(map[TypeKey]TypeValue)	:Membuat map baru
			delete(map, key)			:Menghapus data di map dengan key
	*/

	/// cara lain membuat map
	var book map[string]string = make(map[string]string)

	book["tittle"] = "Belajar Go-lang"
	book["penulis"] = " Fauzi"
	book["ups"] = "salah"

	fmt.Println("Map book :", book)
	fmt.Println("Panjang Map book :", len(book))

	/// mengahpus key map
	delete(book, "ups")
	fmt.Println("Map book setelah menghapus 1 key:", book)

}
