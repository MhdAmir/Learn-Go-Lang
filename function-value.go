package main

import "fmt"

func sayHello(name string) string {
	return "hallo " + name
}
func main() {
	hello := sayHello

	//basic function
	result := sayHello("R")
	fmt.Println(result)

	//function value
	fmt.Println(hello("amir"))
}
