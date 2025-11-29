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

	var r, g, b int
	table := make([][]int, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if j == m-1 {
				fmt.Fscanln(reader, &r, &g, &b)
			} else {
				fmt.Fscan(reader, &r, &g, &b)
			}
			table[i][j] = 2126*r + 7152*g + 722*b
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if table[i][j] < 510000 {
				fmt.Fprint(writer, string(rune(35)))
			} else if table[i][j] < 1020000 {
				fmt.Fprint(writer, string(rune(111)))
			} else if table[i][j] < 1530000 {
				fmt.Fprint(writer, string(rune(43)))
			} else if table[i][j] < 2040000 {
				fmt.Fprint(writer, string(rune(45)))
			} else {
				fmt.Fprint(writer, string(rune(46)))
			}
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
