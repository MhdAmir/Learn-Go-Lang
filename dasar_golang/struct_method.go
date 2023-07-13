/**
struct method
struct = tipe data yang bisa digunakan sebagai parameter untuk fungsi
namun jika ditambahkan method seakan akan struct memiliki function
method = fungsi
*/

package main

import "fmt"

type User struct {
	Name, E_Adress string
	Age            int
}

func (user User) ohayo() {
	fmt.Printf("ohayo %s ^^", user.Name)
}

func main() {
	Amir := User{Name: "Muhammad Amir Abdurrozaq"}
	Amir.ohayo()
}
