/*
TIPE DATA NUMBER PADA GO
1. Integer => tipe data bil bulat.
	Integer dibagi menjadi beberapa jenis tergantung dengan kebutuhannya.

	A.Integer yang bisa diisi dengan nilai negatif
	1. int8
		min : -128
		max : 127
	2. int16
		min : -32.768
		max : 32.767
	3. int32
		min : -2.147.483.648
		max : 2.147.483.647
	4. int64
		min : -9223372036854775808
		max : 9.223.372.036.854.775.807

	B.Integer yang tidak bisa diisi dengan nilai negatif
	1. unint8
		min : 0
		max : 127 * 2
	2. unint16
		min : 0
		max : 32.767 * 2
	3. unint32
		min : 0
		max : 2.147.483.647 * 2
	4. unint64
		min : 0
		max : 9.223.372.036.854.775.807 * 2

	unint (unsigned integer) memiliki 2x nilai maksimum dibanding dengan int biasa , karena
	tidak memiliki nilai negatif.
	semakin besar batas max , maka semakin besar pula memory penyimpanan yang digunakan.

2. Floating Point => tipe data decimal.
	1. float32
		min : 1.18 * 10^-38 (decimal negatif)
		max : 3.4 * 10^38 (decimal positif)
	2. float64
		min : 2.23 * 10^-308 (decimal negatif)
		max : 1.80 * 10^308 (decimal positif)
	3. complex64 : complex number with float32 real dan imaginary parts.
	4. complex128 : complex number with float64 real dan imaginary parts.


ALIAS
kata ganti untuk memanggil type data
1. byte => ia akan memanggil type data unint8.
2. rune => ia akan memanggil type daya int32.
3. int => ia akan memanggil int32 / int64 (sesuai dengan system operasi yang digunakan)
4. unint => ia akan memanggil unint32 / unint64 (sesuai dengan system operasi yang digunakan)
*/

package main

import "fmt"

func main() {
	fmt.Println("Angka satu : ", 1)
	fmt.Println("Angka dua : ", 2)
	fmt.Println("Angka tiga koma lima : ", 3.5)
}
