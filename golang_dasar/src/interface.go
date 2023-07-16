package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Printf("Hello %s\n", hasName.GetName())
}

/**
Implementation Interface
Every data type who compatible(sesuai) with contract interface, automaticly considered as that interface
So that we dont have to implemetation manually

*/

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var Amir Person
	Amir.Name = "Amir"

	beruang := Animal{Name: "Madu"}

	SayHello(Amir)
	SayHello(beruang)
}
