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

	var q, x, t int
	table := []int{}
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &q, &x, &t)
		if q == 2 {
			if len(table) == x {
				table = append(table, t)
			} else {
				table = append(table[:x], append([]int{t}, table[x:]...)...)
			}
		} else {
			if x == 1 {
				sort.Slice(table, func(k, l int) bool {
					return table[k] < table[l]
				})
			} else {
				sort.Slice(table, func(k, l int) bool {
					return table[k] > table[l]
				})
			}
		}
	}

	fmt.Fprintln(writer, len(table))
	for i := 0; i < len(table); i++ {
		fmt.Fprint(writer, table[i], " ")
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
