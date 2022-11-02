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
		table[i] = make([]int, 2)
		fmt.Fscanln(reader, &table[i][0], &table[i][1])
	}

	sort.Slice(table, func(i, j int) bool {
		if table[i][1] == table[j][1] {
			return table[i][0] < table[j][0]
		}
		return table[i][1] < table[j][1]
	})

	for _, t := range table {
		fmt.Fprintln(writer, t[0], t[1])
	}
	writer.Flush()
}
