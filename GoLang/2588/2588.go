package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var num1 int
	var num2 string
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)

	slice := strings.Split(num2, "")

	for i := len(slice) - 1; i >= 0; i-- {
		a, _ := strconv.Atoi(slice[i])
		fmt.Println(num1 * a)
	}

	a, _ := strconv.Atoi(num2)
	fmt.Println(num1 * a)
}
