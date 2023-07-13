/**
type error interface
name of error interface is error

> tidak perlu membuat manual
> golang menyediakan library error untuk memudahkan developer
*/

package main

import (
	"errors"
	"fmt"
)

func divide(numbers int, divider int) (int, error) {
	if divider == 0 {
		//membuat massage baru error dengan errors.New("isi msg")
		return 0, errors.New("Can't Divide with zero")
	} else {
		return numbers / divider, nil
	}
}

func main() {
	// var contohErr error = errors.New("Error gan")
	// fmt.Println(contohErr.Error())

	result, err := divide(10, 0)
	if err == nil {
		fmt.Println(result)

	} else {
		fmt.Println(err.Error())

	}
}
