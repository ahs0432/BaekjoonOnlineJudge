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

	table := []int{}
	for i := 1; i <= n; i++ {
		table = append(table, i)
	}

	var t1, t2 int
	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &t1, &t2)
		table[t1-1], table[t2-1] = table[t2-1], table[t1-1]
	}
	for i := 0; i < n; i++ {
		fmt.Fprint(writer, table[i], " ")
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
