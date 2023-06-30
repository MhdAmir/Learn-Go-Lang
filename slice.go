package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"November",
		"Desember",
	}

	var slice = months[4:7]
	// slice[0] = "MEIKU"
	fmt.Println(slice)
	fmt.Println(months)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	var slice2 = months[2:4]
	fmt.Println(slice2)

	var slice3 = append(slice2, "eko")
	fmt.Println(slice3)
	slice3[1] = "amir"
	fmt.Println(slice2)
	fmt.Println(months)

	copySlice := make([]string, len(slice2), cap(slice2))
	copy(copySlice, slice2)
	fmt.Println(copySlice)

	//this how define array
	iniArray := [...]int{1, 2, 3, 4, 5}
	//this how define slice
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
