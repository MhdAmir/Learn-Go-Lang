/*
*
Pass by Value
> Defaultly in all variable in go-lang is passing by value, not by reference
> mean, that we send variable to the function, method or other variable,
> Actually what is sent is duplicate value
*/

/*
*
Pointer
have ability to make reference to data location in same memory without
duplicating existing data
with Pointer we can make PASS BY REFERENCE
*/
package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// PASS BY VALUE

	// addr1 := Address{"Cicurug", "Sukabumi", "Jawa Barat"}
	// //copied to address2 in different location
	// addr2 := address1

	// addr2.City = "Cibubur"

	// fmt.Println(address1)
	// fmt.Println(address2)

	//PASS BY REFERENCE
	//Operator &
	// addr1 := Address{"Cicurug", "Sukabumi", "Jawa Barat"}
	// //use &var when we input (& meaning is the location address memory some variable)
	// addr2 := &addr1

	// addr2.City = "bandung"
	// //when we change addr2. indirectly we change addr1 too
	// fmt.Println(addr1)
	// fmt.Println(*addr2)

	//Operator * (1)
	// addr1 := Address{"Cicurug", "Sukabumi", "Jawa Barat"}
	// //use &var when we input (& meaning is the location address memory some variable)
	// addr2 := &addr1

	// addr2.City = "bandung"
	// //when we change addr2. indirectly we change addr1 too

	// addr2 = &Address{"Cibubur", "Bekasi", "Jawa Barat"}
	// //addr2 will make different location with addr1 so addr2 didnt change addr1

	// fmt.Println(addr1)
	// fmt.Println(addr2)

	//Operator * (2)
	addr1 := Address{"Cicurug", "Sukabumi", "Jawa Barat"}
	//use &var when we input (& meaning is the location address memory some variable)
	addr2 := &addr1

	addr2.City = "bandung"
	//when we change addr2. indirectly we change addr1 too

	*addr2 = Address{"Cibubur", "Bekasi", "Jakarta"}
	//with *addr2 it will change addr1 & addr2 value because still have same memory address
	//it will make addr1 & addr2 will change into new address and forget old address

	fmt.Println(addr1)
	fmt.Println(addr2)

	//FUNCTION NEW
	/**
	golang have new function who can use for make pointer
	but function new only return pointer into empty data, meaning there is not first data */

	addr4 := new(Address)
	addr5 := addr4

	addr5.City = "Ciawi"
	fmt.Println(addr4)
	fmt.Println(addr5)

}
