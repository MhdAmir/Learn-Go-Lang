package database

import "fmt"

var conn string

func init() {
	fmt.Println("init dipanggil")
	conn = "MySQL"
}

func GetDataBase() string {
	return conn
}
