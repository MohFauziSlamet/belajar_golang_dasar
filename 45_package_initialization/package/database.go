package database

import "fmt"

var connection string

func init() {
	fmt.Println("init dijalankan")
	connection = "MySQL"
	fmt.Println("Success connect to ", connection)
}

func GetDatabase() string {
	return connection
}
