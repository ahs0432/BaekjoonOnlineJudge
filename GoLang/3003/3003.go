package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3, num4, num5, num6 int
	fmt.Scanln(&num1, &num2, &num3, &num4, &num5, &num6)
	fmt.Println((1 - num1), (1 - num2), (2 - num3), (2 - num4), (2 - num5), (8 - num6))
}
