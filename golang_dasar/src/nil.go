package main

import "fmt"

func NewMap(user string) map[string]string {
	if user == "" {
		return nil

	} else {
		return map[string]string{"name": user}
	}
}
func main() {
	user := NewMap("")
	if user == nil {
		fmt.Println("User Kosong")
	} else {
		fmt.Println(user)
	}
}
