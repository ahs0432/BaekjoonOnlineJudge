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

	var n, m int64
	var sum int64
	for i := 1; i <= t; i++ {
		fmt.Fscanln(reader, &n, &m)

		sum = (m - n + 1) * (n + m) / 2
		fmt.Fprintf(writer, "Scenario #%d:\n", i)
		fmt.Fprintln(writer, sum)
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
