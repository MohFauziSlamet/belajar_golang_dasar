package main

import "fmt"

/*
Struct
● Struct adalah sebuah template data yang digunakan untuk
menggabungkan satu atau lebih tipe data lainnya dalam
satu kesatuan.
● Struct biasanya representasi data dalam program
aplikasi yang kita buat.
● Data di struct disimpan dalam field.
● Sederhananya struct adalah kumpulan dari field.
● Dev golang, biasanya membuat nama struct dengan Upercase
contoh : DataCustomer.
● Dan juga Dev golang, biasanya membuat variable/field
didalam struct dengan Upercase
contoh : FirstName / LastName.
*/

/*
Membuat Data Struct
● Struct adalah template data atau prototype data.
● Struct tidak bisa langsung digunakan.
● Namun kita bisa membuat data/object (exp:variabel) ,
dari struct yang telah kita buat.
*/

// Membuat struct
type Customer struct {
	Name, Addres string // penulisan singkat

	// Penulisan panjang
	/*
		Name string
		Addres string
	*/

	Age int
}

/*
Struct Method
● Struct adalah tipe data seperti tipe data lainnya,
dia bisa digunakan sebagai parameter untuk function.
● Namun jika kita ingin menambahkan method ke dalam structs,
sehingga seakan-akan sebuah struct memiliki function.
● Method adalah function.
*/

// contoh func biasa
func sayHi(customer Customer, name string) {
	fmt.Println("Hii", name, "My name", customer.Name)
}

// contoh struct func
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name", customer.Name)
}

func main() {
	fmt.Println()

	fmt.Println("Membuat struct 1")

	var ludy Customer

	// Memasukan data
	ludy.Name = "Barochatul Mauludy"
	ludy.Addres = "Suko"
	ludy.Age = 24

	// Func biasa
	sayHi(ludy, "Fauzi")

	fmt.Println("======================================================================")
	fmt.Println()

	// struct func
	ludy.sayHello("Fauzi")

}
