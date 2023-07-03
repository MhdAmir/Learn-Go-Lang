package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}
func main() {
	//argument statement
	total := sumAll(2, 5, 7, 9)
	fmt.Println(total)

	//slice statement
	slice := []int{10, 10, 10, 10, 10}
	total = sumAll(slice...)
	fmt.Println(total)
}
