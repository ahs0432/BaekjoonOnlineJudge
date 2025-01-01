package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var n, sum int
	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &n)

		sum = 0
		for j := 1; j <= n/2; j++ {
			if n%j == 0 {
				sum += j
			}
		}

		if sum == n {
			fmt.Fprintln(writer, "Perfect")
		} else if sum < n {
			fmt.Fprintln(writer, "Deficient")
		} else {
			fmt.Fprintln(writer, "Abundant")
		}
	}
	writer.Flush()
}
