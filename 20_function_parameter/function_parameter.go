package main

import "fmt"

func sayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	/*
		Function Parameter
		Saat membuat function, kadang-kadang kita membutuhkan data dari luar,
		atau kita sebut parameter.
		Kita bisa menambahkan parameter di function, bisa lebih dari satu.
		Parameter tidaklah wajib, jadi kita bisa membuat function tanpa parameter
		seperti sebelumnya yang sudah kita buat.
		Namun jika kita menambahkan parameter di function,
		maka ketika memanggil function tersebut, kita wajib memasukkan data ke parameternya.
	*/

	/// Memanggil func secara langsung
	sayHello("Moh", "Fauzi")

	/*
		Ketika kita membuat function dengan parameter, saat memanggilnya penempatan parameter
		wajib berurutan. dan type datanya juga harus sesuai.
		jumlah parameter yang diisi harus sesuai, jika lebih akan terjadi error.
	*/

}
