//panic function is function who can we use for stop program
//panic function often use in error while the program run
//when panic function called, defer function still executed

package main

import "fmt"

func endApp() {
	fmt.Println("Apk END")
}
func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
	fmt.Println("Apk run")
}
func main() {
	runApp(true)
}
