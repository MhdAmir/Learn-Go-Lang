package main

import "fmt"

func sayilyto(name string, many int16) {
	fmt.Println(name, "ily", many)
}
func main() {
	name := "re"
	many := 3000
	sayilyto(name, int16(many))
}
