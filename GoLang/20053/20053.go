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

	for i := 0; i < n; i++ {
		var t int
		fmt.Fscanln(reader, &t)

		min := 1000000
		max := -1000000

		for j := 0; j < t; j++ {
			var tmp int

			if j == t-1 {
				fmt.Fscanln(reader, &tmp)
			} else {
				fmt.Fscan(reader, &tmp)
			}

			if tmp < min {
				min = tmp
			}

			if tmp > max {
				max = tmp
			}
		}

		fmt.Println(min, max)
	}
}
