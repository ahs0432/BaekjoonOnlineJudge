package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(reader, &n, &m)

	table := make([][]int, n+1)
	for i := range table {
		table[i] = []int{-1, -1}
	}

	var a, c int
	var b string
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a, &b, &c)

		if b == "P" {
			table[a][0] = c
		} else {
			table[a][1] = c
		}
	}

	minCount, maxCount := 0, 0
	for i := 1; i <= n; i++ {
		if table[i][0] != 0 && table[i][1] != 1 {
			maxCount++
			if table[i][0] == 1 && table[i][1] == 0 {
				minCount++
			}
		}
	}

	fmt.Println(minCount, maxCount)
}
