package main

import (
	"belajar_golang_dasar/43_package_and_import/counting"
	"belajar_golang_dasar/43_package_and_import/helper"
	"fmt"
)

// struct
type Man struct {
	Name string
}

func main() {
	fmt.Println()
	res := helper.SayHello("Barochatul  1")
	resGoodBy := helper.SayGoodBy("Barochatul Mauludy 2")
	counting := counting.Counting("Barochatul Mauludy 3")
	fmt.Println("======================================================================")
	fmt.Println(res)
	fmt.Println(resGoodBy)
	fmt.Println(counting)

}
