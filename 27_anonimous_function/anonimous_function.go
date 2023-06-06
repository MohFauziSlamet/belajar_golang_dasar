package main

import "fmt"

/*
Anonymous Function
● Sebelumnya setiap membuat function, kita akan selalu memberikan
sebuah nama pada function tersebut.
● Namun kadang ada kalanya lebih mudah membuat function secara langsung
di variable atau parameter tanpa harus membuat function terlebih dahulu.
● Hal tersebut dinamakan anonymous function, atau function tanpa nama.

*/

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You are blocked : ", name)
		return
	}

	fmt.Println("Welcome : ", name)

}

func main() {
	fmt.Println()

	fmt.Println("Anonimous Function dengan func di variabel")
	blackList := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blackList)
	registerUser("root", blackList)

	fmt.Println("==============================================================================")
	fmt.Println()

	fmt.Println("Anonimous Function langsung di parameter")

	registerUser("admin", func(s string) bool { return s != "admin" })
	registerUser("admin", func(s string) bool { return s != "root" })
	registerUser("fauzi", func(s string) bool { return s == " fauzi" })

}
