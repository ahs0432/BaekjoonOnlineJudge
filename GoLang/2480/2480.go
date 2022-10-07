package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3 int
	fmt.Scanln(&num1, &num2, &num3)

	if num1 == num2 && num2 == num3 {
		fmt.Print(10000 + (num1 * 1000))
	} else if num1 == num2 || num2 == num3 {
		fmt.Print(1000 + (num2 * 100))
	} else if num1 == num3 {
		fmt.Print(1000 + (num1 * 100))
	} else {
		if num1 > num2 {
			if num1 > num3 {
				fmt.Print(num1 * 100)
			} else {
				fmt.Print(num3 * 100)
			}
		} else {
			if num2 > num3 {
				fmt.Print(num2 * 100)
			} else {
				fmt.Print(num3 * 100)
			}
		}
	}
}
