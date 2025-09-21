package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	table := make([][]int, 3)
	for i := 0; i < 3; i++ {
		table[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &table[0][i], &table[1][i], &table[2][i])
	}

	var counts map[int]int
	score := make([]int, n)
	for i := 0; i < 3; i++ {
		counts = make(map[int]int)
		for j := 0; j < n; j++ {
			counts[table[i][j]]++
		}

		for j := 0; j < n; j++ {
			if counts[table[i][j]] == 1 {
				score[j] += table[i][j]
			}
		}
	}

	for _, s := range score {
		fmt.Fprintln(writer, s)
	}
	writer.Flush()
}
