package main

import "fmt"

/*
Function as Parameter.

Function tidak hanya bisa kita simpan di dalam variable sebagai value.
Namun juga bisa kita gunakan sebagai parameter untuk function lain.
*/

func sayHeloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func filteredName(name string) string {
	/// conditional return
	if name == "Anjing" {
		return "..."
	}

	/// default return
	return name
}

/// ==============================================================================

/*
Function Type Declarations

Kadang jika function terlalu panjang, agak ribet untuk menuliskannya di dalam parameter.
Type Declarations juga bisa digunakan untuk membuat alias function, sehingga akan mempermudah
kita menggunakan function sebagai parameter.
*/
type filter func(string) string

func sayHeloWithFilter2(name string, filter filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func filteredName2(name string) string {
	/// conditional return
	if name == "Anjing" {
		return "..."
	}

	/// default return
	return name
}

func main() {
	fmt.Println()

	fmt.Println("Func as paramater")
	sayHeloWithFilter("Fauzi", filteredName)
	fmt.Println("==============================================================================")
	fmt.Println()

	fmt.Println("Func as paramater with variabel")
	filtered := filteredName
	sayHeloWithFilter("Anjing", filtered)
	fmt.Println("==============================================================================")
	fmt.Println("==============================================================================")
	fmt.Println()

	fmt.Println("Func as paramater with Function Type Declarations")
	sayHeloWithFilter2("Fauzi", filteredName)
	fmt.Println("==============================================================================")
	fmt.Println()

	fmt.Println("Func as paramater with variabel & Function Type Declarations")
	filtered2 := filteredName2
	sayHeloWithFilter("Anjing", filtered2)

}
