package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	table := make([]int, n)

	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Fscanln(reader, &table[i])
		} else {
			fmt.Fscan(reader, &table[i])
		}
	}

	now := 0
	for i := 0; i < m; i++ {
		now += table[i]
	}

	max := now
	for i := m; i < n; i++ {
		now -= table[i-m]
		now += table[i]

		if max < now {
			max = now
		}
	}
	fmt.Println(max)
}
