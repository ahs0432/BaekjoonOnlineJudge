package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	u1, u2 := 100, 100

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscanln(reader, &a, &b)

		if a > b {
			u2 -= a
		} else if a < b {
			u1 -= b
		}
	}

	fmt.Println(u1)
	fmt.Println(u2)
}
