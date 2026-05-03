package main

import (
	"fmt"
)

func main() {
	var num1 int
	fmt.Scanln(&num1)

	if num1 >= 90 {
		fmt.Print("A")
	} else if num1 >= 80 {
		fmt.Print("B")
	} else if num1 >= 70 {
		fmt.Print("C")
	} else if num1 >= 60 {
		fmt.Print("D")
	} else {
		fmt.Print("F")
	}
}
