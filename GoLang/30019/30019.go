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

	table := make([][]int, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, 2)
		table[i][0], table[i][1] = 0, 0
	}

	var k, s, e int
	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &k, &s, &e)
		if table[k-1][1] > s {
			fmt.Fprintln(writer, "NO")
		} else {
			table[k-1][0], table[k-1][1] = s, e
			fmt.Fprintln(writer, "YES")
		}
	}
	writer.Flush()
}
