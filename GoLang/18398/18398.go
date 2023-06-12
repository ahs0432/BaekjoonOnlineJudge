package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t, n int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n)
		for j := 0; j < n; j++ {
			var a, b int
			fmt.Fscanln(reader, &a, &b)
			fmt.Fprintln(writer, a+b, a*b)
		}
	}

	writer.Flush()
}
