/**
Type Ascertions
==> ability to change data type to become desired data type something like convertion
this features often be use when we found with empty data interface
*/

package main

import "fmt"

func random() interface{} {
	return true
}

func main() {
	result := random()
	// resultString := result.(string)

	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("value", value, "-> String")
	case int:
		fmt.Println("value", value, "-> integer")
	default:
		fmt.Println("value", value, "-> Unknown")
	}

}
