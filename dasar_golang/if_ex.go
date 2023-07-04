package main

import "fmt"

func main() {

	name := "rv"

	if name == "amir" {
		fmt.Println("hallo amir")
	} else if name == "rv" {
		fmt.Println("hallo rv")
	} else {
		fmt.Println("hallo, boleh kenal?")
	}

	//short statement
	if n := len(name); n > 5 {
		fmt.Println("terlalu panjang bang")
	} else {
		fmt.Println("nama sudah benar")
	}
}
