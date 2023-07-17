package helper

import "fmt"

//access modifier
//jika huruf diawali dengan huruf kecil
//tidak bisa di pakai di luar package
//error

var jkt = 48
var App = "learn_golang"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodBye(name string) {
	fmt.Println("Bye", name)
}
