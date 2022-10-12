package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	fmt.Scanln(&num1, &num2)

	if num2-45 < 0 {
		num2 = 60 + (num2 - 45)

		if num1 == 0 {
			num1 = 23
		} else {
			num1 = num1 - 1
		}
	} else {
		num2 = num2 - 45
	}

	fmt.Print(num1, num2)
}
