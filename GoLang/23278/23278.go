package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, l, h int
	fmt.Fscanln(reader, &n, &l, &h)

	table := make([]float64, n)
	sum := float64(0)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &table[i])
		sum += table[i]
	}

	sort.Slice(table, func(i, j int) bool {
		return table[i] < table[j]
	})

	for i := 0; i < l; i++ {
		sum -= table[i]
	}
	for i := 0; i < h; i++ {
		sum -= table[n-i-1]
	}
	fmt.Println(sum / float64(n-l-h))
}
