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

	table := make([]int, 3)
	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &table[0], &table[1], &table[2])

		sort.Slice(table, func(a, b int) bool {
			return table[a] < table[b]
		})
		fmt.Fprintf(writer, "Case #%d: ", i)

		if table[0]*table[0]+table[1]*table[1] == table[2]*table[2] {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
	writer.Flush()
}
