package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	table := make([][]int, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, 2)
		fmt.Fscanln(reader, &table[i][0], &table[i][1])
	}

	var max, tmp, pi, pj int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			tmp = abs(table[i][0]-table[j][0]) + abs(table[i][1]-table[j][1])
			if tmp > max {
				max = tmp
				pi = i
				pj = j
			}
		}
	}
	fmt.Println(pi+1, pj+1)
}
