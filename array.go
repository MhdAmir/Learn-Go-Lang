package main

import "fmt"

func main() {
	var name [3]string
	name[0] = "Muhammad"
	name[1] = "Amir"
	name[2] = "Abdurrozaq"
	// name[3] = "woy"  it will be error. because there is only 3 defined array the index only until n-1

	fmt.Println(name)
	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])
	// fmt.Println(name[3]) it will be error. because there is only 3 defined array the index only until n-1

	//more simple define
	var values = [10]int16{
		0,
		1,
		3,
		5,
		7,
		9,
		//each other value who undefined have 0 value
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(values[3])
	fmt.Println(values[4])
	fmt.Println(values[5])
	fmt.Println(values[6])
	fmt.Println(values[7])
	fmt.Println(values[8])
	fmt.Println(values[9])

	fmt.Println(len(name[0]))

}
