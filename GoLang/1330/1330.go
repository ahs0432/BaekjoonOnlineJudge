package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	fmt.Scanln(&num1, &num2)

	if num1 > num2 {
		fmt.Print(">")
	} else if num1 < num2 {
		fmt.Print("<")
	} else {
		fmt.Print("==")
	}
}
