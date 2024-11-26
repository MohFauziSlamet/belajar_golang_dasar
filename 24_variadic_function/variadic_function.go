package main

import "fmt"

// func variadic, hanya bisa di letakkan di akhir parameter.
// jika diletakan di tengah atau didepan, akan error.
// variabel argumen hanya ada satu pada setiap function.
func sumAll(firstName string, lastName string, numbers ...int) (string, int) {
	name := firstName + lastName
	total := 0
	for _, number := range numbers {
		total += number
	}
	return name, total

}

func main() {
	/*
		Variadic Function
		Parameter yang berada di posisi terakhir,
		memiliki kemampuan dijadikan sebuah var-args (var argumen).
		Var-args artinya datanya bisa menerima lebih dari satu input,
		atau anggap saja semacam Array / slice.
		Apa bedanya dengan parameter biasa dengan tipe data Array?
		Jika parameter tipe Array, kita wajib membuat array
		terlebih dahulu sebelum mengirimkan ke function.
		Jika parameter menggunakan var-args, kita bisa
		langsung mengirim data nya,
		jika lebih dari satu, cukup gunakan tanda koma.

	*/

	name1, number1 := sumAll("Moh Fauzi ", "Slamet", 1, 2, 3, 4, 5, 6)

	fmt.Println(name1, number1)

	fmt.Println("==============================================================================")

	/*
		Jika data yang akan dimasukan kedalam parameter berupa slice,
		maka ketika mengisi pada parameter wajib di ikuti titik 3x.
	*/

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	name2, number2 := sumAll("Moh Fauzi ", "Slamet", slice...)

	fmt.Println(name2, number2)

	fmt.Println("==============================================================================")

}
