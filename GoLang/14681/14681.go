package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)

	if num1 > 0 && num2 > 0 {
		fmt.Printf("1")
	} else if num1 < 0 && num2 > 0 {
		fmt.Printf("2")
	} else if num1 < 0 && num2 < 0 {
		fmt.Printf("3")
	} else if num1 > 0 && num2 < 0 {
		fmt.Printf("4")
	}
}
