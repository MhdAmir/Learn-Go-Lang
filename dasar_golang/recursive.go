package main

import "fmt"

func factorial(nilai int) int {
	if nilai > 1 {
		return nilai * factorial(nilai-1)
	} else {
		return 1
	}

}
func main() {
	fmt.Println(factorial(5))
}
