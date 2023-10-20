package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	max := int64(9223372036854775807)
	min := int64(-9223372036854775808)

	var n, of int
	var tmp, sum int64
	for i := 0; i < 3; i++ {
		sum, of = 0, 0
		fmt.Fscanln(reader, &n)
		for j := 0; j < n; j++ {
			fmt.Fscanln(reader, &tmp)

			if sum > 0 && tmp > 0 && tmp > max-sum {
				of += 1
			}

			if sum < 0 && tmp < 0 && tmp < min-sum {
				of -= 1
			}

			sum += tmp
		}

		if of > 0 {
			fmt.Fprintln(writer, "+")
		} else if of < 0 {
			fmt.Fprintln(writer, "-")
		} else if sum > 0 {
			fmt.Fprintln(writer, "+")
		} else if sum < 0 {
			fmt.Fprintln(writer, "-")
		} else {
			fmt.Fprintln(writer, 0)
		}
	}

	writer.Flush()
}
