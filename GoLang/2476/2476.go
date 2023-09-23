package main

import (
	"bufio"
	"fmt"
	"os"
)

func max_int(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	max := 0

	var a, b, c int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b, &c)

		sum := 0
		if a == b && b == c {
			sum = 10000 + (1000 * a)
		} else if a == b || a == c {
			sum = 1000 + (100 * a)
		} else if b == c {
			sum = 1000 + (100 * b)
		} else {
			sum = 100 * (max_int(a, max_int(b, c)))
		}

		if sum > max {
			max = sum
		}
	}
	fmt.Println(max)
}
