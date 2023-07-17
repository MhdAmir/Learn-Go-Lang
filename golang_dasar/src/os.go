/**
Package OS
=> fungsionalitas untuk akses fitur OS secara independen
ALL OS FRIENDLY*/

package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)

	fmt.Printf("%s %s %s\n", args[1], args[2], args[3])

	//untuk mengetahuai hostname
	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", name)
	} else {
		fmt.Println("Error ", err.Error())
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}
