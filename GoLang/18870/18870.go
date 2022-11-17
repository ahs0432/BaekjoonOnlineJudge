package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	table := make([][]int, n)

	for i := 0; i < n; i++ {
		table[i] = make([]int, 3)
		if i == n-1 {
			fmt.Fscanln(reader, &table[i][0])
		} else {
			fmt.Fscan(reader, &table[i][0])
		}
		table[i][1] = i
	}

	sort.Slice(table, func(i, j int) bool {
		return table[i][0] < table[j][0]
	})

	table[0][2] = 0
	for i := 1; i < n; i++ {
		table[i][2] = table[i-1][2]
		if table[i][0] != table[i-1][0] {
			table[i][2]++
		}
	}

	sort.Slice(table, func(i, j int) bool {
		return table[i][1] < table[j][1]
	})

	for i := 0; i < n; i++ {
		fmt.Fprint(writer, table[i][2], " ")
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
