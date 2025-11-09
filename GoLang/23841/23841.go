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

	table := make([][]byte, n)
	for i := 0; i < n; i++ {
		table[i] = make([]byte, m)
		fmt.Fscanln(reader, &table[i])

		for j := 0; j < m/2; j++ {
			if table[i][j] == '.' {
				table[i][j] = table[i][m-j-1]
			} else {
				table[i][m-j-1] = table[i][j]
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, string(table[i]))
	}
	writer.Flush()
}
