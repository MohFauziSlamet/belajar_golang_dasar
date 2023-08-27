package main

///* dengan memanggil salah satu function pada package.
// import (
// 	database "belajar_golang_dasar/45_package_initialization/package"
// 	"fmt"
// )

///* tanpa memanggil function apapun, hanya memanggil init pada package
///  Blank Identifier
import (
	_ "belajar_golang_dasar/45_package_initialization/package"
)

/*
Package Initialization
● Saat kita membuat package, kita bisa membuat sebuah
function yang akan diakses ketika package
kita diakses.
● Ini sangat cocok ketika contohnya,
jika package kita berisi function-function untuk
berkomunikasi dengan database, kita membuat function
inisialisasi untuk membuka koneksi ke database.
● Untuk membuat function yang diakses secara otomatis
ketika package diakses, kita cukup membuat function
dengan nama init.

*/

/*
Blank Identifier
● Kadang kita hanya ingin menjalankan init function di package tanpa harus mengeksekusi salah satu
function yang ada di package
● Secara default, Go-Lang akan komplen ketika ada package yang di import namun tidak digunakan
● Untuk menangani hal tersebut, kita bisa menggunakan blank identifier (_) sebelum nama package
ketika melakukan import
*/
func main() {
	// dengan memanggil salah satu function pada package.
	// fmt.Println(database.GetDatabase())
}
