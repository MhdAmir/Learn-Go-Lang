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
}
