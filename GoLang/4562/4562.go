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

		if a < b {
			fmt.Fprintln(writer, "NO BRAINS")
		} else {
			fmt.Fprintln(writer, "MMM BRAINS")
		}
	}

	writer.Flush()
}
