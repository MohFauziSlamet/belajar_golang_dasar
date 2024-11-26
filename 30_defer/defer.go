package main

import "fmt"

/*
Defer
● Defer function adalah function yang bisa
kita jadwalkan untuk dieksekusi setelah sebuah
function selesai di eksekusi.
● Defer function akan selalu dieksekusi
walaupun terjadi error di function yang dieksekusi.
● Defer function disarankan diletakkan di awal
sebuah function, apabila diletakan di akhir atau di tengah,
jika terjadi error maka defer bisa saja
tidak dijalankan , tergantung letak baris error tsb.
jika error lebih dulu daripada defer, maka defer tidak dijalankan.
*/

// Func for defer
func succesLogging() {
	fmt.Println("Selesai memanggil function")
}

func runAplication(value int) {
	defer succesLogging()

	// akan terjadi error , jika value dimasukan 0.
	result := 10 / value
	fmt.Println("Run Aplication: ", result)

	/*
		meskipun terjadi error, func defer akan tetap dijalankan.
	*/

}

func main() {
	fmt.Println()

	runAplication(1)
	runAplication(0)

}
