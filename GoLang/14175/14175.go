package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var m, n, k int
	fmt.Fscanln(reader, &m, &n, &k)

	table := make([]string, m)
	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &table[i])
	}

	var now string
	for i := 0; i < m; i++ {
		now = ""
		for j := 0; j < n; j++ {
			for l := 0; l < k; l++ {
				now += string(table[i][j])
			}
		}

		for l := 0; l < k; l++ {
			fmt.Fprintln(writer, now)
		}
	}
	writer.Flush()
}
