package main

import "fmt"

// SAME AS C/C++
type Address struct {
	City, Province, Country string
}

//but the pointer in data type
func changeCountryToSurabaya(alamat *Address) {
	alamat.Country = "Surabaya"
}
func main() {
	alamat := Address{"Cicurug", "Sukabumi", "Jawa Barat"}
	fmt.Println(alamat)
	changeCountryToSurabaya(&alamat)
	fmt.Println(alamat)
}
