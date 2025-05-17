package main

import (
	"bufio"
	"fmt"
	"os"
)

func draw(table [][]int, s, n int) {
	if n <= 0 {
		return
	}

	for i := s; i < s+4*(n-1)+1; i++ {
		table[s][i] = 1
		table[i][s] = 1
		table[s+4*(n-1)][i] = 1
		table[i][s+4*(n-1)] = 1
	}

	draw(table, s+2, n-1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	table := make([][]int, 4*(n-1)+1)

	for i := 0; i < len(table); i++ {
		table[i] = make([]int, 4*(n-1)+1)
	}

	draw(table, 0, n)
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table); j++ {
			if table[i][j] == 1 {
				fmt.Fprint(writer, "*")
			} else {
				fmt.Fprint(writer, " ")
			}
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
