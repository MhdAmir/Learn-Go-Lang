package main

import "fmt"

func getAllDate() (string, int16, int8) {
	return "july", 2023, 1
}

func main() {
	month, year, date := getAllDate()
	fmt.Println(month)
	fmt.Println(year)
	fmt.Println(date)

}
