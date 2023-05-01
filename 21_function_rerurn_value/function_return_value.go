package main

import "fmt"

func sayHello(name string) string {

	if name == "" {
		return "Hello bro"
	}

	return "Hello " + name

}

func main() {
	/*
		Function Return Value
		Function bisa mengembalikan data.
		Untuk memberitahu bahwa function mengembalikan data,
		kita harus menuliskan tipe data kembalian dari function tersebut.
		Jika function tersebut kita deklarasikan dengan tipe data pengembalian,
		maka wajib di dalam function nya kita harus mengembalikan data.
		Untuk mengembalikan data dari function, kita bisa menggunakan kata kunci return,
		diikuti dengan datanya.

	*/

	fmt.Println(sayHello("Fauzi"))
	fmt.Println(sayHello(""))

	/// Dengan ditampung di dalam variabel
	res := sayHello("Mauludy")
	fmt.Println(res)

}
