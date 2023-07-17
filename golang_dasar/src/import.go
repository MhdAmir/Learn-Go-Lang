package main

import (
	"fmt"
	"golang_dasar/helper"
)

func main() {
	helper.SayHello("amir")
	fmt.Println(helper.App)

	// helper.sayGoodBye("masa lalu") ERROR
	// fmt.Println(helper.jkt)  ERROR
}
