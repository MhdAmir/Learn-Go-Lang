package main

import "fmt"

func main() {
	var number32 int32 = 129
	var number64 int64 = int64(number32)
	var number8 int8 = int8(number32)

	fmt.Println("number 16 : ", number32)
	fmt.Println("number 32 : ", number64)
	fmt.Println("number 8 : ", number8)

	var name = "Amir"
	var a = name[0]
	var aString = string(a)

	fmt.Println(name)
	fmt.Println(aString)
}
