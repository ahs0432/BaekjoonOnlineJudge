package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	table := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &table[i])
	}

	var x, total, yCnt, nCnt int
	for i := 1; i < n; i++ {
		total = 0
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &x)
			total += abs(table[j] - x)
		}

		if total > 2000 {
			yCnt++
		} else {
			nCnt++
		}
	}

	if yCnt >= nCnt {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
