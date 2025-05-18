package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	table := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &table[i])
	}
	sort.Slice(table, func(i, j int) bool {
		return table[i] < table[j]
	})

	fmt.Printf("%.2f\n", table[1])
}
