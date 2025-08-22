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

	var s, n, q, p int

	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &s)
		fmt.Fscanln(reader, &n)

		for j := 0; j < n; j++ {
			fmt.Fscanln(reader, &q, &p)
			s += (q * p)
		}

		fmt.Fprintln(writer, s)
	}
	writer.Flush()
}
