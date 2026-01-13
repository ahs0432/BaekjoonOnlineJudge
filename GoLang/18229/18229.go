package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m, k int
	fmt.Fscanln(reader, &n, &m, &k)

	table := make([][]int, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if j == m-1 {
				fmt.Fscanln(reader, &table[i][j])
			} else {
				fmt.Fscan(reader, &table[i][j])
			}
		}
	}

	check := make([]int, n)
	for i := 0; ; i++ {
		for j := 0; j < n; j++ {
			check[j] += table[j][i]
			if check[j] >= k {
				fmt.Println(j+1, i+1)
				return
			}
		}
	}
}
