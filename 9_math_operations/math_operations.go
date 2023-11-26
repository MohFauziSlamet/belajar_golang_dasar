package main

import "fmt"

func main() {
	/*
		Operasi Matematika
		Operator + Keterangan Penjumlahan
		Operator - Keterangan Pengurangan
		Operator * Keterangan Perkalian
		Operator / Keterangan Pembagian
		Operator % Keterangan Sisa pembagian (modulo)
	*/

	var a = 10
	var b = 10
	var c = a + b
	var d = a - b
	var e = a * b
	var f = a / b
	var g = a % b

	println(c)
	println(d)
	println(e)
	println(f)
	println(g)

	fmt.Println("==============================================================================")

	/*
		Augmented Assignments
		Operasi Matematika : a = a + 10 == Augmented Assignments : a += 10
		Operasi Matematika : a = a - 10 == Augmented Assignments : a -= 10
		Operasi Matematika : a = a * 10 == Augmented Assignments : a *= 10
		Operasi Matematika : a = a / 10 == Augmented Assignments : a /= 10
		Operasi Matematika : a = a % 10 == Augmented Assignments : a %= 10
	*/

	var i = 10

	i += 10
	fmt.Println("i : ", i)

	i -= 10 // i = 20
	fmt.Println(i)

	i *= 10 // i = 10
	fmt.Println(i)

	i /= 10 // i = 100
	fmt.Println(i)

	i %= 10 // i = 10
	fmt.Println(i)

	fmt.Println("==============================================================================")

	/*
		Unary Operator
		Operator ++	Keterangan a = a + 1
		Operator --	Keterangan a = a - 1
		Operator -	Keterangan negative
		Operator +	Keterangan positive
		Operator !	Keterangan boolean kebalikan (neggasi)
	*/

	a = 10
	a++
	println(a)

	a-- // a = 11
	println(a)

	var negative = -100
	var positive = +100 // bisa langsung di input 100 , tanda + tidak wajib
	println(negative)
	println(positive)

	var isMerried = true

	println(!isMerried)
}
