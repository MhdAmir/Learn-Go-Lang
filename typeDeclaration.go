package main

import "fmt"

func main() {
	type Name string
	type Online bool

	var player1 Name = "Darth Vader"
	var statusP1 Online = true

	fmt.Println("Name of Player 1 : ", player1)
	fmt.Println("Status : ", statusP1)

}
