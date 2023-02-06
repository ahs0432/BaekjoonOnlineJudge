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

	table := make([]int, m)

	check := 1
	now := 1
	for i := 0; i < m; i++ {
		now--
		table[i] = check

		if now == 0 {
			check++
			now = check
		}
	}

	sum := 0
	for i := n - 1; i < m; i++ {
		sum += table[i]
	}
	fmt.Println(sum)
}
