package main

import "fmt"

func main() {
	var x = 10
	var y = 40

	var z float32

	z = float32(x) / float32(y)

	fmt.Println("z = ", z)
}
