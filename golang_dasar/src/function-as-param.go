package main

import "fmt"

//function type declaration
type FILTER func(string) string

//function not only can save in variable
//but can to be parameters other function

func sayHelloWithFilter(name string, filter FILTER) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Babi" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Amir", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Babi", filter)
}
