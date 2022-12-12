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

	table := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &table[i])
	}

	sort.Slice(table, func(i, j int) bool {
		return table[i] < table[j]
	})

	for _, t := range table {
		fmt.Fprintln(writer, t)
	}
	writer.Flush()
}
