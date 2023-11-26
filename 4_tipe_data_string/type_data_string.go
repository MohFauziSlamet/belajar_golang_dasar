/*
TIPE DATA STRING PADA GO
 - String adalah type data kumpulan karakter.
 - Jumlah karakter pada string bisa 0 sampai tak hingga.
 - type data string pada go di tuliskan dengan kata kunci string.
 - nilai string di awali dengan kutip dua , dan diakhiri dengan
   kutip dua ("").


 FUNCTION ynag bisa digunakan didalam string
 len("string") => menghitung jumlah karakter yang ada di string tsb.
 "string"[index] => mengambil karakter pada posisi yang telah ditentukan
 (ia akan menampilkan memory lokasi byte nya , bukan string value nya.)

*/

package main

import "fmt"

func main() {
	fmt.Println("MOH")
	fmt.Println("FAUZI ")
	fmt.Println("SLAMET")
	fmt.Println("=========")
	fmt.Println(len("MOH"))
	fmt.Println("FAUZI"[0])
	fmt.Println(string("FAUZI"[0]))
	fmt.Println("SLAMET"[1])         // untuk melihat lokasi byte nya
	fmt.Println(string("SLAMET"[1])) // untuk melihat value , dari lokasi byte tsb.
}
