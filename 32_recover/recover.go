package main

import "fmt"

/*
Recover
● Recover adalah function yang bisa kita gunakan untuk
menangkap data panic.
● Dengan recover proses panic akan terhenti,
sehingga program akan tetap berjalan.
● Recover wajib di masukan kedalam func, yang menjadi defer.


*/

// Func for defer
func endApp() {
	fmt.Println("End App")

	message := recover()

	fmt.Println("error message : ", message)
}

func runAplication(error bool) {
	defer endApp()

	if error {
		panic("error")
	}

	fmt.Println("Run Aplication")

}

func main() {
	fmt.Println()

	runAplication(true)

	// cek , jika error di func atas,
	// ini akan tetap dijalankan
	fmt.Println("Ludy")

}
