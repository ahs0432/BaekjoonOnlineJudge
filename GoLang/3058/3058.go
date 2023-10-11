package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var tmp, min, sum int

	for i := 0; i < n; i++ {
		min, sum = 101, 0
		for j := 0; j < 7; j++ {
			fmt.Fscan(reader, &tmp)

			if tmp%2 == 0 {
				sum += tmp

				if tmp < min {
					min = tmp
				}
			}
		}
		fmt.Fprintln(writer, sum, min)
	}

	writer.Flush()
}
