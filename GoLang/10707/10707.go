package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c, d, p int
	x, y := 0, 0
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)
	fmt.Fscanln(reader, &c)
	fmt.Fscanln(reader, &d)
	fmt.Fscanln(reader, &p)

	x = a * p
	if c < p {
		y = b + ((p - c) * d)
	} else {
		y = b
	}

	if x < y {
		fmt.Println(x)
	} else {
		fmt.Println(y)
	}
}
