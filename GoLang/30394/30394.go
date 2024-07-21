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

	var x, y int64
	min, max := int64(9223372036854775807), int64(-9223372036854775808)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &x, &y)

		if y < min {
			min = y
		}

		if y > max {
			max = y
		}
	}

	fmt.Println(max - min)
}
