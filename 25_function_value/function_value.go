package main

import "fmt"

func sayHello(name string) string {

	return "Hay " + name

}

func main() {
	/*
		Function Value
		Function adalah first class citizen.
		Function juga merupakan tipe data, dan bisa disimpan di dalam variable.

	*/

	/// Bisa di tampung dulu ke dalam variabel , atau bisa juga langsung di panggil di print
	sayHello := sayHello
	hello := sayHello("Mauludy")
	fmt.Println(hello)
	fmt.Println("==============================================================================")
	fmt.Println(sayHello("Fauzi"))

}
