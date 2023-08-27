package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println()

	/*
		Package os
		● Go-Lang telah menyediakan banyak sekali package bawaan, salah satunya adalah package os
		● Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen (bisa
		digunakan  disemua sistem operasi)
		● https://golang.org/pkg/os/

	*/

	fmt.Println("Argument")
	fmt.Println(os.Args)

	fmt.Println("Hostname")
	name, err := os.Hostname()
	if err == nil {
		fmt.Println(name)
	} else {
		fmt.Println(err)

	}

	fmt.Println("ENV")
	fmt.Println("Username", os.Getenv("USER"))
	fmt.Println("Password", os.Getenv("password"))

}
