package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	table := make([][]int, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &table[i][j])
		}
	}

	store1 := []int{}
	store2 := []int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if table[i][j] == 1 {
				if len(store1) == 0 {
					store1 = []int{i + 1, j + 1}
				} else {
					store2 = []int{i + 1, j + 1}
				}
			}
		}
	}

	fmt.Println(abs(store2[0]-store1[0]) + abs(store2[1]-store1[1]))
}
