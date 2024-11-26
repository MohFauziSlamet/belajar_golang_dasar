package main

import "fmt"

func main() {
	/*
		Break & Continue
		Break & continue adalah kata kunci yang
		bisa digunakan dalam perulangan.
		Break digunakan untuk menghentikan seluruh perulangan.
		sedangkan
		Continue adalah digunakan untuk menghentikan
		perulangan yang saat ini berjalan,
		dan langsung melanjutkan ke perulangan selanjutnya,
		jika ada post statement maka
		akan dijalankan post statement.
	*/

	fmt.Println("Break")
	for i := 0; i <= 10; i++ {
		fmt.Println("Index ke :", i)
		if i == 5 {
			break
		}
	}

	fmt.Println("======================================================================")
	fmt.Println()

	fmt.Println("Continue")
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Index ke :", i)
	}

}
