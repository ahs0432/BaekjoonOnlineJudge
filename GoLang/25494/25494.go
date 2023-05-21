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

	for i := 0; i < n; i++ {
		var a, b, c int
		fmt.Fscanln(reader, &a, &b, &c)

		if a > b {
			a = b
		}

		if a > c {
			a = c
		}

		fmt.Fprintln(writer, a)
	}
	writer.Flush()
}
