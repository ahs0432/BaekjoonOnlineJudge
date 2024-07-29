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

	table := make([]int, n+1)
	for i := 1; i <= n; i++ {
		table[i] = i
	}

	var a, b int
	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &a, &b)

		for j := 0; j <= (b-a)/2; j++ {
			table[a+j], table[b-j] = table[b-j], table[a+j]
		}
	}

	for i := 1; i <= n; i++ {
		fmt.Fprint(writer, table[i], " ")
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
