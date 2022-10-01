package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3 int
	fmt.Scanln(&num1, &num2, &num3)
	fmt.Println((num1 + num2) % num3)
	fmt.Println(((num1 % num3) + (num2 % num3)) % num3)
	fmt.Println((num1 * num2) % num3)
	fmt.Println(((num1 % num3) * (num2 % num3)) % num3)
}
