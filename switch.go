package main

import "fmt"

func main() {
	name := "amir"
	// name := "Amir"
	// name := "AMIR"

	switch name {
	case "amir":
		fmt.Println("name huruf kecil semua")
	case "AMIR":
		fmt.Println("name huruf besar semua")
	case "Amir":
		fmt.Println("name sudah benar")
	}

	//switch with short statement
	switch lenght := len(name); lenght > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah sesuai")
	}

	//switch without condition
	switch {
	case len(name) == 5:
		fmt.Println("nama 5 kata")

	case len(name) > 5:
		fmt.Println("nama lebih dari 5 kata")

	case len(name) < 5:
		fmt.Println("nama kurang dari 5 kata")

	}
}
