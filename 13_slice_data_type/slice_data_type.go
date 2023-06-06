package main

import "fmt"

func main() {
	/*
				Tipe Data Slice
				Tipe data Slice adalah potongan dari data Array.
				Slice mirip dengan Array, yang membedakan adalah ukuran Slice bisa berubah.
				Slice dan Array selalu terkoneksi, dimana Slice adalah data yang mengakses sebagian atau seluruh data di Array


				Detail Tipe Data Slice
				Tipe Data Slice memiliki 3 data, yaitu pointer, length dan capacity.
				Pointer adalah penunjuk data pertama di array pada slice
				Length adalah panjang dari slice, dan
				Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity.

				Ketentuan Membuat Slice Dari Array
		 		Membuat Slice				Keterangan
				1. slice[low:high]			Membuat slice dari array dimulai index low sampai index sebelum high
				2. slice[low:]				Membuat slide dari array dimulai index low sampai index akhir di array
				3. slice[:high]				Membuat slice dari array dimulai index 0 sampai index sebelum high
				4. slice[:]					Membuat slice dari array dimulai index 0 sampai index akhir di array

				Function Slice
				 Operasi							Keterangan
				1. len(slice)Untuk 					mendapatkan panjang
				2. cap(slice)Untuk 					mendapat kapasitas
				3. append(slice, data)				Membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array baru
				4. make([]TypeData, length, capacity)			Membuat slice baru
				5. copy(destination, source)					Menyalin slice dari source ke destination

	*/

	/// Membuat array bulan
	month := [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"Nobember",
		"Desember",
	}

	/// Membuat slice dari array [month]
	/// Membuat dari index ke 4 sampai ke 6 . Batas exclusive 7.
	slice1 := month[4:7]

	fmt.Println("Data slice :", slice1, "Panjang lenght :", len(slice1), "Panjang capasity :", cap(slice1)) // [Mei Juni Juli] Panjang lenght : 3 Panjang capasity : 8

	/*
		Hati hati jika menggunakan slice.
		Jika array parent[month] di ubah , maka data slice [slice1] akan juga ikut berubah,
		dan sebaliknya,
		jika slice [slice1] di ubah , maka array parent [month] juga akan ikut berubah.

		jadi kesimpulannya array dan slice saling me reference (terhubung).
	*/

	/// Contoh mengubah data array
	// month[5] = "Ubah"
	// fmt.Println("Data array yang di ubah pada index ke 5", month)
	// fmt.Println("Data slice akan juga ikut berubah", slice1)

	/// Contoh mengubah slice
	// slice1[1] = "Ubah"
	// fmt.Println("Data slice dibuah pada index ke 1 :", slice1)
	// fmt.Println("Data array juga akan ikut berubah :", month)

	slice2 := month[10:]
	fmt.Println("Slice2 :", slice2) //  [Nobember Desember]

	/// membuat append dari slice 2
	slice3 := append(slice2, "Fauzi")
	fmt.Println("Slice3 :", slice3) //  [Nobember Desember Fauzi]

	/// kita ubah data slice3 pada index tertentu
	slice3[1] = "Bukan Desember"

	/// kita coba cek semua data, mana yang berubah
	fmt.Println("Slice 3 : ", slice3)     //  [Nobember Bukan Desember Fauzi]
	fmt.Println("Slice 2 : ", slice2)     //  [Nobember Desember]
	fmt.Println("Array parent : ", month) //  [Januari Februari Maret April Mei Juni Juli Agustus September Oktober Nobember Desember]

	/*
		Jika kita membuat append slice3, ada dua kemungkinan yang akan terjadi.
		jika capasity slice2  yang mau kita jadikan append, masih bisa menampung.
		maka data akan dijadikan satu dengan array parent yaitu [month] , karena [month] adl sumber
		dari [slice2].


		apabila capasity slice2 array tidak mencukupi,
		maka akan di buat slice array baru.

		dalam 2 hal ini memiliki dampak yang berbeda.

		jika masih dalam satu array parent [month], ketika kita mengubah [slice3],
		maka [slice2] dan [month] akan ikut berubah, karena saling reference.

		jika dalam array yang berbeda, ketika kita mengubah [slice3], data
		[slice2] dan [month] tidak akan berubah. karena sudah array yang berbeda.

		contoh didalam array yang sama pada [slice4]
	*/

	/// contoh append , dalam satu array.
	/// karena capacity [slice1] adalah 8,
	/// dan panjang [slice4] setelah di append adalah 4, maka
	/// akan di tampung pada array sama sama yaitu [momth] (saling berkaitan / reference).
	slice4 := append(slice1, "Fauzi")

	fmt.Println()
	fmt.Println("Slice4 : ", slice4)            //  [Mei Juni Juli Fauzi]
	fmt.Println("Slice1 : ", slice1)            //  [Mei Juni Juli Fauzi]
	fmt.Println("Array parent month : ", month) //  "Agustus" berubah jadi "Fauzi" => [Januari Februari Maret April Mei Juni Juli Fauzi September Oktober Nobember Desember]
	/*
	 ketika membuat append [slice4] , dan [slice1] masih bisa menampung capasity nya,
	 maka array parent[month] juga akan berdampak.
	 index yang berubah pada array parent [month], bergantung pada
	 index di [slice1] ketika pembuatan slice.

	 misal month memiliki 12 index.
	 month [0,1,2,3,4,5,6,7,8,9,10,11]
	 dan [slice1] di ambil dari index[4,5,6] => memiliki capasity 8 (index 4 sampai index 11)

	 ketika membuat append dari [slice1] yang memiliki capasity 8, dan masih bisa menampung.
	 misal append slice4 := [slice1,"Fauzi"] , ini akan menempati index [4,5,6,7] pada array parent [month]

	 maka secara langsung, value pada month[7] akan berubah menjadi "Fauzi"
	*/

	/// kita coba ubah [slice4]
	/// maka [slice1] dan [month] akan berubah
	slice4[1] = "Mauludy"

	/// kita coba cek [slice4] [slice1] [month]
	fmt.Println("Cek slice4 : ", slice4) // juni berubah menjadi mauludy =>  [Mei Mauludy Juli Fauzi]
	fmt.Println("Cek slice1 : ", slice1) // juni berubah menjadi mauludy =>  [Mei Mauludy Juli]
	fmt.Println("Cek month : ", month)   // juni berubah menjadi mauludy =>   [Januari Februari Maret April Mei Mauludy Juli Fauzi September Oktober Nobember Desember]

	fmt.Println("===============================")

	/*
		Membuat slice dari awal , tanpa reference ke array yang lain dengan func make()
		hal ini lebih aman, karena tidak menggantu array lain.
	*/

	var newSlice = make([]string, 2, 5)
	newSlice[0] = "Fauzi"
	newSlice[1] = "Mauludy"

	fmt.Println("newSlice : ", newSlice)

	/*
		Ketika melakukan copy slice, pastikan panjang dan capasity harus sama,
		jika tidak, maka data akan terpotong.
	*/
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println("copySlice : ", copySlice)

	fmt.Println("===============================")

	/*
		Cara lain membuat slice
		Hati-Hati Saat Membuat Array
		Saat membuat Array, kita harus berhati-hati, jika salah, maka yang kita buat bukanlah Array, melainkan Slice
	*/

	/// bisa memakai panjang, jika kita tau panjang yang akan di tampung berapa dengan menuliskan [5]
	/// atau
	/// jika kita tidak tau berapa panjang yang akan ditampung, bisa menuliskan [...]
	var iniArray = [5]int{1, 2, 3, 4, 5}

	/// tidak menuliskan panjang
	var iniSlice = []int{1, 2, 3, 4, 5} // tidak memakai panjang

	fmt.Println("INI ARRAY : ", iniArray)
	fmt.Println("INI SLICE : ", iniSlice)
}
