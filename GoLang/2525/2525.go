package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3 int
	fmt.Scanln(&num1, &num2)
	fmt.Scanln(&num3)

	num2 = num2 + int(num3%60)
	num1 = (num1 + int(num3/60) + int(num2/60)) % 24
	num2 = (num2 % 60)

	fmt.Print(num1, num2)
}
