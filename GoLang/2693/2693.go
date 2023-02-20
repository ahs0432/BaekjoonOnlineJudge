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

	for i := 0; i < n; i++ {
		table := make([]int, 10)
		for j := 0; j < 10; j++ {
			if j == 9 {
				fmt.Fscanln(reader, &table[j])
			} else {
				fmt.Fscan(reader, &table[j])
			}
		}
		sort.Slice(table, func(i2, j int) bool {
			return table[i2] > table[j]
		})

		fmt.Fprintln(writer, table[2])
	}
	writer.Flush()
}
