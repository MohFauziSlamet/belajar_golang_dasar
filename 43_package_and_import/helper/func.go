/*
penamaan package, setidaknya harus sama dengan nama
dari foldernya.

Package
● Package adalah tempat yang bisa digunakan untuk mengorganisir kode program yang kita buat di
Go-Lang
● Dengan menggunakan package, kita bisa merapikan kode program yang kita buat
● Package sendiri sebenarnya hanya direktori folder di sistem operasi kita

Ketentuan membuat package di golang.
● didalam satu folder package tidak bisa terdapat 2 nama function
yang sama. jadi penamaan function harus di bedakan.
● Jika didalam folder berbeda, maka di bolehkan menggunakan
nama function yang sama dengan package lain.
● Jika mengakses didalam nama file yang sama, maka tinggal
menuliskan namanya langsung. misal nama func / variable.
● Jika mengakses diluar nama file yang sama, maka harus
melakukan import terlebih dahulu , setelah itu menuliskan
nama packagenya di ikuti nama func / variable .
misal nama func / variable.
*/
// package helper

// func SayHello(name string) string {
// 	return "Hello " + name
// }

package helper

import "fmt"

func SayGoodBy(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
