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
	names := make([]string, n)

	for i := 0; i < n; i++ {
		table[i] = make([]int, 2)
		fmt.Fscanln(reader, &table[i][0], &names[i])
		table[i][1] = i
	}

	sort.Slice(table, func(i int, j int) bool {
		if table[i][0] == table[j][0] {
			return table[i][1] < table[j][1]
		}
		return table[i][0] < table[j][0]
	})

	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, table[i][0], names[table[i][1]])
	}

	writer.Flush()
}
