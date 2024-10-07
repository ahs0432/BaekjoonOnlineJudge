package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t, n, tmp int
	fmt.Fscanln(reader, &t)

	var min, max int
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n)

		min, max = 100, -1
		for j := 0; j < n; j++ {
			if j == n-1 {
				fmt.Fscanln(reader, &tmp)
			} else {
				fmt.Fscan(reader, &tmp)
			}

			if min > tmp {
				min = tmp
			}

			if max < tmp {
				max = tmp
			}
		}
		fmt.Fprintln(writer, (max-min)*2)
	}
	writer.Flush()
}
