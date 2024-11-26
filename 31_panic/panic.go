package main

import "fmt"

/*
Panic
● Panic function adalah function yang bisa kita gunakan
untuk menghentikan program.
● Panic function biasanya dipanggil ketika terjadi error
pada saat program kita berjalan.
● Saat panic function dipanggil, program akan terhenti,
namun defer function tetap akan dieksekusi.

*/

// Func for defer
func endApp() {
	fmt.Println("End App")
	recoverMess := recover()
	fmt.Println(recoverMess, "\n\n============================")
}

func runAplication(error bool) {
	defer endApp()

	if error {
		panic("Error")
	}

	fmt.Println("Run Aplication")

}

func main() {
	fmt.Println()

	runAplication(true)
	runAplication(false)

}
