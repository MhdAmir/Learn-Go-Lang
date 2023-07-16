package main

import "fmt"

/**
method akan disimpan dalam struct,
tapi sebenernya masih pass by value*/

// type Man struct {
// 	Name string
// }

// func (man Man) Married() {
// 	man.Name = "Mr. " + man.Name
// }

// func main() {
// 	pria := Man{"Budi"}
// 	pria.Married()

// 	fmt.Printf(pria.Name)
// }

// OUTPUT ==> Budi
/**
jadi yang di ubah dalam method nilainya dimasukkan dalam alamat yang berbeda
sehingga tidak akan mengubah nilai yang sudah ada di dalam fungsi main*/

//cukup menambah * di method

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	pria := Man{"Budi"}
	pria.Married()

	fmt.Printf(pria.Name)
}
