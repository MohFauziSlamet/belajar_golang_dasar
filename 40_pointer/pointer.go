package main

import (
	"fmt"
)

// struct(1)
type Address struct {
	city, province, country string
}

func main() {
	fmt.Println()

	/*
		Pass by Value
		● Secara default di Go-Lang semua variable itu di passing by value,
		bukan by reference.
		● Artinya, jika kita mengirim sebuah variable ke dalam function,
		method atau variable lain, sebenarnya yang dikirim adalah
		duplikasi value nya (value1 di copy ke value2).
		● dalam kasus ini , kita akan membuat contoh pada struct.(1)
	*/

	address1 := Address{city: "Malang", province: "Jawa timur", country: "Indonesia"}

	// membuat copy dari address1
	address2 := address1

	// kita tampilkan semua
	fmt.Println("Addres 1 :", address1)
	fmt.Println("Addres 2 :", address2)

	fmt.Println("======================================================================")
	fmt.Println()

	fmt.Println("Mengubah salah satu data di addres 2")
	// sekarang kita coba ubah isi data dari addres 2
	address2.city = "Batu"

	// kita tampilkan semua lagi
	fmt.Println("Addres 1 :", address1) // Addres 1 : {Malang Jawa timur Indonesia}
	fmt.Println("Addres 2 :", address2) // Addres 2 : {Batu Jawa timur Indonesia}

	/*
		hasilnya yang berubah hanya addres 2. karena addres dua mengarah
		pada memory yang berbeda dengan address 1. jadi ada 2 lokasi memori
		yaitu untuk address1 dan address2.
	*/

	fmt.Println("======================================================================")
	fmt.Println()

	/*
		Pointer
		● Pointer adalah kemampuan membuat reference ke lokasi data di memory
		yang sama, tanpa menduplikasi data yang sudah ada.
		● Sederhananya, dengan kemampuan pointer, kita bisa membuat
		pass by reference.
		● ini merupakan kebalikan dari pass by value. dan juga lebih menghemat
		penggunaan memori.

		Operator &
		● Untuk membuat sebuah variable dengan nilai pointer ke variable yang
		lain, kita bisa menggunakan operator & diikuti dengan nama variable nya.

	*/

	addressLudy := Address{city: "Malang", province: "Jawa timur", country: "Indonesia"}
	// membuat reference dari addressLudy
	addressFauzi := &addressLudy

	fmt.Println("Menampilkan Addres ludy dan Addres fauzi sebelum diubah")

	fmt.Println("Addres ludy :", addressLudy)
	fmt.Println("Addres fauzi :", addressFauzi)

	fmt.Println()

	// lalu kita ubah data di address fauzi
	addressFauzi.city = "Batu"
	fmt.Println("Menampilkan Addres ludy dan Addres fauzi sesudah diubah")
	// kita tampilkan semua
	fmt.Println("Addres ludy :", addressLudy)
	fmt.Println("Addres fauzi :", addressFauzi)
	fmt.Println()

	/*
		Addres ludy : {Batu Jawa timur Indonesia}
		Addres fauzi : &{Batu Jawa timur Indonesia}
		hasil nya addressLudy akan ikut berubah, jika addressFauzi diubah.
		hal ini karena addressFauzi mengarah pada pointer yang sama dengan
		addressLudy.
	*/

	/*
		Operator *
		● Saat kita mengubah variable,
		maka yang berubah hanya variable tersebut.
		(yaitu meng asign ulang suatu variable dengan value yang baru)
		● Semua variable yang mengacu ke data yang sama tidak akan
		berubah.
		● Jika kita ingin mengubah seluruh variable yang mengacu ke data
		tersebut, kita bisa menggunakan operator *.

		sebelumnya kita membuat sebuah variabel addressFauzi sbg pointer
		dari variable addressLudy.

		kita akan mencoba meng asign ulang addresFauzi.
		apakah addresLudy akan berubah?
	*/

	addressFauzi = &Address{city: "Surabaya", province: "Jawa Timur", country: "Indonesia"}
	/*
		kita telah mengubah variabel addressFauzi.
		kita coba cek semuanya. apakah addressLudy juga ikut berubah?
	*/

	fmt.Println("Menampilkan asign ulang addressFauzi")
	// kita tampilkan semua
	fmt.Println("Addres ludy :", addressLudy)   //  {Batu Jawa timur Indonesia}
	fmt.Println("Addres fauzi :", addressFauzi) //  &{Surabaya Jawa Timur Indonesia}
	fmt.Println()

	/*
		Hasilnya kedua address memiliki value yang bberbeda.
		seperti penjelasan di atas.
		ketika kita melakukan perubahan pada variabel (asign ulang),
		maka pointer tidak akan ikut berubah.

		jika kita ingin mengubah semua data, yang mengacu pada pointer
		yang sama. maka kita bisa menggunakan tanda bintang(*) diawal
		dan sebelum nama variabel yang akan di asign ulang.
	*/

	address3 := Address{city: "JAKARTA", province: "JAWA BARAT", country: "INDONESIA"}
	address4 := &address3

	address4.city = "SEMARANG"

	// mengubah semua data.
	*address4 = Address{city: "Malang", province: "Jawa Timur", country: "Indonesia"}

	address5 := &address3
	fmt.Println("Menampilkan asign ulang addressFauzi dengan tanda bintang")
	// kita tampilkan semua
	fmt.Println("Addres address3 :", address3)
	fmt.Println("Addres address4 :", address4)
	fmt.Println("Addres address5 :", address5)
	fmt.Println()

	/*
		address5 akan ikut berubah
	*/

	fmt.Println("======================================================================")
	fmt.Println()

	/*
		Function new
		● Sebelumnya untuk membuat pointer dengan menggunakan operator &.
		● Go-Lang juga memiliki function new yang bisa digunakan
		untuk membuat pointer.
		● Namun function new hanya mengembalikan pointer ke data kosong,
		artinya tidak ada data awal.
	*/

	address6 := new(Address)
	fmt.Println("Pointer kosong")
	fmt.Println(address6)

	fmt.Println()
	address6.city = "Blitar"
	fmt.Println(address6)

}
