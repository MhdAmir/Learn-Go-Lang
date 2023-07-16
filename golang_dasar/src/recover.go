//recover == function who can use for take panic data
//with recover panic process will stop, so the program will run again

package main

import "fmt"

func endApp() {
	fmt.Println("Apk END")
	// no error code
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message", message)
	}
}
func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
	// main.runApp(0x0?)
	// C:/workspace/Belajar Go-Lang/recover.go:14 +0x10c
	// wrong code
	// still error because when panic run
	// the below program from panic will stop
	// message := recover()
	// fmt.Println("Error dengan message", message)
	fmt.Println("Apk run")
}
func main() {
	runApp(true)
	fmt.Println("re")
}
