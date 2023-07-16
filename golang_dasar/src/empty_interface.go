/**
Empty Interface
=> Interface who have any declaration method, which is make automaticly all data type will be implement
examp:
> fmt.Println(a ...interface{})
> panic(v ...interface{})
> recover() interface{}

*/

package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}
}

func main() {
	//tipe datanya juga harus berisi interface kosong (interface{})
	var data interface{} = Ups(2)
	fmt.Println(data)
}
