package main

import "fmt"

func main() {
	// const YourName = "Muh Amir Abd"
	// const NumberAge = 12
	// const phi = 3.14

	//can be declared like this

	const (
		YourName  = "Muh Amir Abd"
		NumberAge = 12
		phi       = 3.14
	)

	fmt.Println(YourName)
	fmt.Println(NumberAge)
	fmt.Println(phi)

}
