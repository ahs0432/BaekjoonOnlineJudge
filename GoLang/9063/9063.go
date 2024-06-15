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

	var x, y int
	min_x, min_y, max_x, max_y := 10000, 10000, -10000, -10000
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &x, &y)

		if x < min_x {
			min_x = x
		}

		if x > max_x {
			max_x = x
		}

		if y < min_y {
			min_y = y
		}

		if y > max_y {
			max_y = y
		}
	}

	fmt.Println((max_x - min_x) * (max_y - min_y))
}
