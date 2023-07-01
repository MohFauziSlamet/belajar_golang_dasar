package main

import (
	"fmt"
)

// struct
type Man struct {
	Name string
}

// Struct funct tanpa pointer (1)
func (man Man) Married() {
	man.Name = "Mrs. " + man.Name

	// pengecekan nama (tag 2)
	fmt.Println(man.Name)
}

// Struct funct dengan pointer (3)
// cukup menambahkan tanda bintang pada nama type nya.
func (man *Man) MarriedPointer() {
	man.Name = "Mr. " + man.Name
}

func main() {
	fmt.Println()

	/*
		Pointer di struct funtion.
		● Walaupun method akan menempel di struct,
		tapi sebenarnya data struct yang diakses di dalam
		method adalah pass by value (dicopy).
		● Sangat direkomendasikan menggunakan pointer di method,
		sehingga tidak boros memory karena harus selalu
		diduplikasi ketika memanggil method.

	*/

	// contoh struct func biasa (1)
	ludy := Man{Name: "Mauludy"}

	ludy.Married()

	/*
		Apakah akan bisa merubah data variable asli.
	*/
	fmt.Println(ludy)

	/*
		Hasilnya, variable asli tidak berubah, dan menghasilkan
		output {Mauludy} , tanpa di awali dengan Mr.

		kita coba cek, apakah didalam func,
		nama sudah berubah.( tag 2).

		ketika dijalankan, dialam function nama sudah berubah :
		Mr. Mauludy.

		namun di luar function, nama masih Mauludy.


		hal ini terjadi karena, saat memanggil function,
		data yang dikirim hanya di copy kedalam func ,
		bukan data asli.
	*/

	fmt.Println("======================================================================")
	fmt.Println()

	// contoh struct func biasa (1)
	fauzi := Man{Name: "Fauzi"}

	fauzi.MarriedPointer()

	/*
		Apakah akan bisa merubah data variable asli.
	*/
	fmt.Println(fauzi)

	/*
		Hasilnya, variable asli berubah, dan menghasilkan
		output {Mr. Fauzi}.

		hal ini terjadi karena, saat memanggil function,
		data asli yang dikirim kedalam func. sehingga nama yang
		di inisiasi diawal bisa berubah didalam func.
	*/

}
