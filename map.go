package main

import "fmt"

func main() {
	person := map[string]string{
		"name":       "Amir",
		"department": "Informatics Engineer",
	}

	kelas := map[float32]string{
		0.12: "King",
		3.14: "phi",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["department"])

	fmt.Println(kelas)
	fmt.Println(kelas[0.1])
	fmt.Println(kelas[3.14])

	var book map[string]string = make(map[string]string)

	book["title"] = "belajar mengenai masa depan"
	book["price"] = "20k"
	book["author"] = "M Amir A"
	fmt.Println(book)

	delete(book, "title")
	fmt.Println(book)

}
