//defer function == function which can be dated for executioned after
//other function finished
//always executed no matter error

package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil fungsi")
}

func runApp(value int) {
	defer logging()
	result := 10 / value
	fmt.Println("result", result)
	fmt.Println("run App")
	// panic: runtime error: integer divide by zero
	// logging()
}
func main() {
	runApp(0)

}
