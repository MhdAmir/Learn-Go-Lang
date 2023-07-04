package main

import "fmt"

//harap menggunakan closure dengan bijak karena jika salah
//maka scope luar fungsi bisa masuk nilainya ke dalam fungsi

func main() {
	name := "Amir"
	counter := 0

	increment := func() {
		name = "re"
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
