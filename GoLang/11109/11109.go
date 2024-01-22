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

	var d, n, s, p int
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &d, &n, &s, &p)
		if d+p*n == n*s {
			fmt.Fprintln(writer, "does not matter")
		} else if d+p*n > n*s {
			fmt.Fprintln(writer, "do not parallelize")
		} else {
			fmt.Fprintln(writer, "parallelize")
		}
	}
	writer.Flush()
}
