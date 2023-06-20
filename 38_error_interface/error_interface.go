package main

import (
	"errors"
	"fmt"
)

/*
error Interface
● Go-Lang memiliki interface yang digunakan sebagai kontrak untuk membuat error,
nama interface nya adalah error.

Membuat Error
● Untuk membuat error, kita tidak perlu manual.
● Go-Lang sudah menyediakan library untuk membuat helper secara mudah,
yang terdapat di package errors.
*/

func Pembagian(nilai int, pembagi int) (int, error) {

	if pembagi == 0 {
		return 0, errors.New("Pembagi Tidak boleh nol")
	}

	// Default
	return nilai / pembagi, nil

}

func main() {
	fmt.Println()

	value, err := Pembagian(10, 0)

	fmt.Println(value, err)

	fmt.Println("======================================================================")
	fmt.Println()

}
