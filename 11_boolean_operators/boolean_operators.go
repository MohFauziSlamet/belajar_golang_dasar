package main

func main() {
	/*
		Operasi Boolean
		Operator		Keterangan
		&&				Dan
		|| 				Atau
		!				Kebalikan(negasi)

		Operasi && (DAN)
		Nilai 1 	Operator	Nilai 2			Hasil
		true			&& 		true			true
		true			&& 		false			false
		false			&& 		true			false
		false			&& 		false			false

		kesimpulannya , ketika semua nilai bernilai true
		maka hasilnya true.

		Operasi || (OR)
		Nilai 1 	Operator	Nilai 2			Hasil
		true			|| 		true			true
		true			|| 		false			true
		false			|| 		true			true
		false			|| 		false			false

		kesimpulannya , ketika salah satu nilai bernilai true
		maka hasilnya true.

		Operasi ! (negasi)
		Nilai 		Operator			Hasil
		true			! 				false
		false			! 				true
	*/

	var nilaiAkhir = 80
	var nilaiAbsensi = 70

	// step by step
	var lulusNilaiAkhir = nilaiAkhir >= 75
	var lulusNilaiAbsensi = nilaiAbsensi >= 75

	var lulus = lulusNilaiAkhir && lulusNilaiAbsensi

	println(lulus)

	// atau dengan cara ini
	println(nilaiAkhir >= 75 && nilaiAbsensi >= 75)

}
