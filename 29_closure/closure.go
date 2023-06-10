package main

import "fmt"

/*
Closures
● Closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam
scope yang sama.
● Harap gunakan fitur closure ini dengan bijak saat kita membuat aplikasi.

Ketika kita salah dalam menyikapi apa itu closure , kita bisa saja
mengubah data yang seharus nya tidak dirubah.
*/

func main() {
	fmt.Println()

	counter := 0

	increment := func() {
		fmt.Println("increment")

		// counter++ adl mengubah nilai variabel yang di atas func.
		// inilah salah satu bentuk closure.
		// dapat mengakses variabel diatas nya.
		counter++

		/*
		 variable name tidak bisa di aksek di dalam func.
		 karena berada di bawah func.
		*/
		// fmt.Println(name)

		/*
			namun
			fmt.Println(counter)
			bisa di akses di dalam func.
			karena variable berada di atas function
		*/

		/*
			Ketika didalam function, kita tidak mau merubah data di luar func
			, kita harus mendeklarasi kan ulang variable tersebut.

			contoh counter.
			counter adalah variabel milik scope di atas func.
			jika tidak mau merubah variable tsb, maka kita wajib mendklarasikan
			variable lagi.

			artinya variable yang kita akan ubah, tidak akan mengubah variable
			di scope atasnya.
		*/

	}

	// name := "Fauzi" // variabel name di buat di bawa func

	// kita coba panggil variabel func nya
	increment()
	increment()

	fmt.Println()
	fmt.Println(counter)
}
