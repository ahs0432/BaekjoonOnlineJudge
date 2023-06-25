package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	item := make([][]int, n)
	for i := 0; i < n; i++ {
		item[i] = make([]int, 2)
		fmt.Fscanln(reader, &item[i][0], &item[i][1])
	}

	bag := make([][]int, n)
	bag[0] = make([]int, k+1)
	for w := 0; w <= k; w++ {
		if item[0][0] > w {
			bag[0][w] = 0
		} else {
			bag[0][w] = item[0][1]
		}
	}

	for i := 1; i < n; i++ {
		bag[i] = make([]int, k+1)
		for j := 1; j <= k; j++ {
			if item[i][0] > j {
				bag[i][j] = bag[i-1][j]
			} else {
				bag[i][j] = max(bag[i-1][j], bag[i-1][j-item[i][0]]+item[i][1])
			}
		}
	}

	fmt.Println(bag[n-1][k])
}
