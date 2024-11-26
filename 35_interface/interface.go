package main

import "fmt"

/*
Interface
● Interface adalah tipe data Abstract,
dia tidak memiliki implementasi langsung.
● Interface hanya berisikan definisi
method (Function).
● Biasanya interface digunakan sebagai kontrak.
● Interface juga bisa digunakan sebagai
parameter pada sebuah func.
*/

// Membuat interface
type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// ======

/*
Implementasi Interface
● Setiap tipe data yang sesuai
dengan kontrak interface,
secara otomatis dianggap sebagai
interface tersebut.
jadi ketika membuat sebuah struct,
dan nama nya sama dengan interface,
maka secara langsung mengimplementasi
sebuah interface.

● Sehingga kita tidak perlu
mengimplementasikan interface
secara manual.

● Hal ini agak berbeda dengan bahasa
pemrograman lain yang ketika membuat
interface, kita harus menyebutkan
secara eksplisit akan
menggunakan interface mana.
*/

// membuat struct
type Person struct {
	Name string
}

// membuat struct func dengan return Interface
func (person Person) GetName() string {
	return person.Name
}

// ======

// membuat struct yang lain dengan implemtasi interface.
// artinya, interface tak terbatas di implemen berapapun struct.

type Animal struct {
	Name string
}

// membuat struct func dengan return Interface
func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	fmt.Println()

	// var ludy HasName
	// sayHello(ludy)
	/*
		interface tidak bisa digunakan secara langsung, harus di deklarasikan.
		kode di atas akan error ketika di run,
		dan menampilkan pesan error : panic: runtime error: invalid memory address or nil pointer dereference.
	*/

	// contoh penggunaan interface
	person := Person{Name: "Barochatul Mauludy"}

	/*
		ketika struct func di hapus / dikomen. maka sayHello akan menampilkan
		error sbb :
		cannot use person (variable of type Person) as HasName value in
		argument to sayHello: Person does not implement HasName
		(missing method GetName)compilerInvalidIfaceAssign.

		dan kenapa tidak error ketika struct function dibuka.
		alasan nya , karena struct func memiliki return yang sama dengan
		interface. artinya secara gak langsung, struct memiliki kontrak
		dengan interface.
	*/
	sayHello(person)

	fmt.Println("======================================================================")
	fmt.Println()

	// contoh penggunaan interface
	cat := Animal{Name: "Temon"}
	sayHello(cat)

}
