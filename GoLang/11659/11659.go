package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	table := make([]int, n)
	for i := 0; i < n; i++ {
		if n-1 == i {
			fmt.Fscanln(reader, &table[i])
		} else {
			fmt.Fscan(reader, &table[i])
		}

		if i != 0 {
			table[i] += table[i-1]
		}
	}

	for i := 0; i < m; i++ {
		var s, e int
		fmt.Fscanln(reader, &s, &e)
		if s == 1 {
			fmt.Fprintln(writer, table[e-1])
		} else {
			fmt.Fprintln(writer, table[e-1]-table[s-2])
		}
	}
	writer.Flush()
}
