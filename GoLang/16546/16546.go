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

	table := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &table[i])
	}

	sort.Slice(table, func(i, j int) bool {
		return table[i] < table[j]
	})

	for i := 0; i < n-1; i++ {
		if i+1 != table[i] {
			fmt.Println(i + 1)
			return
		}
	}
	fmt.Println(n)
}
