package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println()

	host := flag.String("host", "localhost1", "Put your database host")
	user := flag.String("user", "root1", "Put your database user")
	password := flag.String("password", "root1", "Put your database password")

	flag.Parse()

	fmt.Println("host :", host, "|| user :", user, "|| password :", password)
	fmt.Println("host :", *host, "|| user :", *user, "|| password :", *password)

}
