package main

import (
	"fmt"
)

func main() {
	var num1 int
	fmt.Scanln(&num1)

	if (num1%4 == 0 && num1%100 != 0) || num1%400 == 0 {
		fmt.Printf("1")
	} else {
		fmt.Printf("0")
	}
}
