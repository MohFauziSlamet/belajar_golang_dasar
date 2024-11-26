package main

import "fmt"

/*
Recursive Function
● Recursive function adalah function yang
memanggil function dirinya sendiri.
● Kadang dalam pekerjaan, kita sering
menemui kasus dimana menggunakan
recursive function lebih mudah dibandingkan
tidak menggunakan recursive function.
● Contoh kasus yang lebih mudah diselesaikan
menggunakan recursive adalah Factorial.
*/

// Kita bandingkan dengan looping
// factorial looping
func factorialLoop(value int) int {
	res := 1

	// buat perulangan
	for i := value; i > 0; i-- {
		res *= i // res = res * i
	}

	return res
}

// factorial dengan recursive
func factorialRecursive(value int) int {
	// jika value == 1 , maka langsung return
	if value == 1 {
		return value
	}

	return value * factorialRecursive(value-1)

}

func main() {
	fmt.Println()

	fmt.Println("Factorial dengan looping")
	fmt.Println(factorialLoop(5))
	fmt.Println(5 * 4 * 3 * 2 * 1)

	fmt.Println("==============================================================================")
	fmt.Println()

	fmt.Println("Factorial dengan recursive")
	fmt.Println(factorialRecursive(6))
	fmt.Println(6 * 5 * 4 * 3 * 2 * 1)

}
