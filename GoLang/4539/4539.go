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

	var x int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &x)

		for j := 10; j < x; j *= 10 {
			if x%j/(j/10) >= 5 {
				x = (x/j + 1) * j
			} else {
				x = x / j * j
			}
		}
		fmt.Fprintln(writer, x)
	}
	writer.Flush()
}
