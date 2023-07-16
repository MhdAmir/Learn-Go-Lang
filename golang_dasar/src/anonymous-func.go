package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}
func main() {
	//method 1
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("Amir", blacklist)
	registerUser("admin", blacklist)

	//method 2
	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("admin", func(name string) bool {
		return name == "root"
	})

}
