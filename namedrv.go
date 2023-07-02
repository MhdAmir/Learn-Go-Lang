package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Muhammad"
	middleName = "Amir"
	lastName = "Abdurrozaq"
	return
}
func main() {
	firstName, midName, lastName := getCompleteName()
	fmt.Println(firstName, midName, lastName)
}
