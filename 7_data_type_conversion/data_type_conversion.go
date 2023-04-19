package main

import "fmt"

func main() {
	/*
		Konversi Tipe Data
		Di Go-Lang kadang kita butuh melakukan konversi tipe data dari satu tipe ke tipe lain
		Misal kita ingin mengkonversi tipe data int32 ke int63, dan lain-lain.
	*/

	var int32 int32 = 1000
	var int64 int64 = int64(int32)
	var int8 int8 = int8(int32)

	fmt.Println(int32)
	fmt.Println(int64)
	fmt.Println(int8)

	/*
		jika saat konversi dari int besar ke int kecil, maka akan terjadi penyesuain pada int kecil.
		karena memiliki batas max yang berbeda.
		misal int32 = 1000
		ketika di konversi ke int8 akan bernilai = -24.
		alur terjadi perubahanya, dimulai dari batax max yaitu 127 , lalu
		langsung turun ke batas min -128 , lalu turun lagi -127 , -126 , sampai batas max lagi.
		jika sudah sampai batas max lagi , namun ankgka int32 belum selesai di konversi, maka akan
		turun ke batas min lagi , lanjut seperti tadi prosesnya, sampai int32 di konversi semua.
	*/

	fmt.Println("==============================================================================")

	/*
		Sebelumnya kita telah mempelajari cara pengambilan string berdasarkan indexnya,
		namun hasilnya masih berupa data byte.
		untuk mengubah data byte ke bentuk string, cara nya dengan memasukan perintah
		string(dataByteNya)
	*/
	name := "FAUZI"
	// eByte := name[0] // cara 1
	// eByteString := string(eByte)
	eByteString := string(name[0]) // cara 2

	fmt.Println(name)
	// fmt.Println(eByte)
	fmt.Println(eByteString)

}
