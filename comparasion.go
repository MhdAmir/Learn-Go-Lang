package main

import "fmt"

func main() {
	name1 := "amir"
	name2 := "amir"

	var result bool = name1 == name2
	fmt.Println(result)

	value1 := 3500
	value2 := 100

	fmt.Println("value1 > value2 >>", value1 > value2)
	fmt.Println("value1 == value2 >>", value1 == value2)
	fmt.Println("value1 < value2 >>", value1 < value2)
	fmt.Println("value1 != value2 >>", value1 != value2)

}
