package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, c int
	fmt.Fscanln(reader, &n, &c)

	a, b := n, n
	var x, y int
	for i := 0; i < c; i++ {
		fmt.Fscanln(reader, &x, &y)

		if x >= a || y >= b {
			continue
		}

		if b*x < a*y {
			b = y
		} else {
			a = x
		}
	}
	fmt.Println(a * b)
}
