package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()

	/*
		Package strings
		● Package strings adalah package yang berisikan
		function-function untuk memanipulasi tipe data
		String.
		● Ada banyak sekali function yang bisa kita gunakan.
		● https://golang.org/pkg/strings/

		Beberapa Function di Package strings
		Function 								Kegunaan
		strings.Trim(string, cutset) 			Memotong cutset di awal dan akhir string
		strings.ToLower(string) 				Membuat semua karakter string menjadi lower case
		strings.ToUpper(string) 				Membuat semua karakter string menjadi upper case
		strings.Split(string, separator) 		Memotong string berdasarkan separator
		strings.Contains(string, search) 		Mengecek apakah string mengandung string lain
		strings.ReplaceAll(string, from, to) 	Mengubah semua string dari from ke to

	*/

	fmt.Println("string contains :", strings.Contains("Barochatul Mauludy", "ludy"))
	fmt.Println("string split :", strings.Split("Barochatul Mauludy", " "))
	fmt.Println("string toLower :", strings.ToLower("Barochatul Mauludy"))
	fmt.Println("string toUpper :", strings.ToUpper("Barochatul Mauludy"))
	fmt.Println("string trim :", strings.Trim(" Barochatul Mauludy ", " "))
	fmt.Println("string replaceAll :", strings.ReplaceAll(" Barochatul Mauludy ", "Maulu", "AI"))
}
