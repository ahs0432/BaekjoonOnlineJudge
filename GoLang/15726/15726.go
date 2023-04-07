package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c float64
	fmt.Fscanln(reader, &a, &b, &c)

	if b > c {
		fmt.Println(int(a * b / c))
	} else {
		fmt.Println(int(a / b * c))
	}
}
