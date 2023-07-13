/*
*struct
template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
disimpan dalam sebuah field
struct ==== kumpulan sebuah field
*/
package main

import "fmt"

//biasanya di tulis huruf besar terlebih dahulu

type User struct {
	Name, E_Adress string
	Age            int
}

func main() {
	//cara memasukan data manual
	//CARA I
	var Amir User
	Amir.Name = "Muhammad Amir Abdurrozaq"
	Amir.E_Adress = "m.amir.hs19@gmail.com"
	Amir.Age = 19

	fmt.Println(Amir)
	fmt.Println(Amir.Name)
	fmt.Println(Amir.E_Adress)
	fmt.Println(Amir.Age)

	//CARA II
	Budi := User{
		Name:     "BudiUtomo",
		E_Adress: "budiutomo@gmail.com",
		Age:      32,
	}
	fmt.Println(Budi)

	Rev := User{"Rev", "Rev@gmail.com", 20}
	fmt.Println(Rev)

}
