package main

import (
	"fmt"
)

func main() {
	var num1 int
	fmt.Scanln(&num1)

	sum := 0
	for i := 1; i <= num1; i++ {
		sum += i
	}
	fmt.Print(sum)
}
