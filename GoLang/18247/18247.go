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

	var a, b int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b)

		if a >= 12 && b >= 4 {
			fmt.Fprintln(writer, 11*b+4)
		} else {
			fmt.Fprintln(writer, -1)
		}
	}
	writer.Flush()
}
