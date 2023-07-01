package main

import "fmt"

func main() {
	//normal for loop
	counter := 0
	for counter <= 10 {
		//for loop with statement
		for counterInner := counter; counterInner <= 10; counterInner++ {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
		counter++
	}

	slice := []string{"Muhammad", "Amir", "Abdurrozaq"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//for range array || slice
	for i, value := range slice {
		fmt.Println("Indeks", i, "=", value)
	}

	book := make(map[string]string)
	book["title"] = "Apa iyah?"
	book["price"] = "10k"

	for key, value := range book {
		fmt.Println(key, "=", value)
	}
}
