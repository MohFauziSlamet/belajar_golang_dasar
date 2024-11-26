package main

import "fmt"

/*
Interface Kosong
● Go-Lang bukanlah bahasa pemrograman yang berorientasi objek.
● Biasanya dalam pemrograman berorientasi objek, ada satu data parent
di puncak yang bisa dianggap sebagai semua implementasi data yang ada
di bahasa pemrograman tersebut.
● Contoh di Java ada java.lang.Object
● Untuk menangani kasus seperti ini, di Go-Lang kita bisa menggunakan
interface kosong.
● Interface kosong adalah interface yang tidak memiliki deklarasi method
satupun, hal ini membuat secara otomatis semua tipe data akan
menjadi implementasi nya.


Penggunaan Interface Kosong di Go-Lang
● Ada banyak contoh penggunaan interface kosong di Go-Lang, seperti :
○ fmt.Println(a ...interface{})
○ panic(v interface{})
○ recover() interface{}
○ dan lain-lain
*/

// interface kosong
type Apapun interface{}

// ini sama dengan func dibawa , yang memiliki return interface kosong
/*
func ups (i int) interface {}  {} 	ini
bisa diganti dengan
func ups (i int) Apapun  {} 		ini
*/

func ups(i int) Apapun {
	if i == 1 {
		return 1
	}

	if i == 2 {
		return true
	}

	// default
	return " Ups"
}

func main() {
	fmt.Println()

	// contoh salah
	// walaupun interface kosong bisa menampung data apapun,
	// namun tidak bisa di simpan dalam type data selain interface kosong.
	// var data int = ups(1)
	// fmt.Println(data)
	// akan menampilkan :
	// cannot use ups(1) (value of type Apapun) as int value in variable declaration:
	// need type assertioncompilerIncompatibleAssign

	// contoh benar
	var data1 interface{}
	fmt.Println("data1", data1) // output : data1 <nil>

	fmt.Println("======================================================================")
	fmt.Println()

	var data2 interface{} = ups(1)
	fmt.Println("data2", data2)

	fmt.Println("======================================================================")
	fmt.Println()

	var data3 interface{} = ups(2)
	fmt.Println("data2", data3)

	fmt.Println("======================================================================")
	fmt.Println()

	var data4 interface{} = ups(3)
	fmt.Println("data2", data4)

	fmt.Println("======================================================================")
	fmt.Println()

	result := ups(00)
	if intValue, ok := result.(int); ok {
		fmt.Println("Hasil sebagai int:", intValue)
	} else {
		fmt.Println("Tipe data bukan int")
	}

}
