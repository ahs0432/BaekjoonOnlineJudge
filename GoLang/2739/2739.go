package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int
	fmt.Scanln(&num1)

	for i := 1; i < 10; i++ {
		fmt.Println(strconv.Itoa(num1) + " * " + strconv.Itoa(i) + " = " + strconv.Itoa(num1*i))
	}
}
