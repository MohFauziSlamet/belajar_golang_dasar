package main

/*
	Access Modifier
	● Di bahasa pemrograman lain, biasanya ada kata kunci yang bisa digunakan untuk menentukan
	access modifier terhadap suatu function atau variable
	● Di Go-Lang, untuk menentukan access modifier, cukup dengan nama function atau variable
	● Jika nama nya diawali dengan hurup besar, maka artinya bisa diakses dari package lain, jika dimulai
	dengan hurup kecil, artinya tidak bisa diakses dari package lain
*/

var version = "1.2.3"     // tidak bisa di akses dari luar
var Aplication = "Golang" // bisa di akses dari luar

func main() {

}
