package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, i int
	fmt.Fscanln(reader, &n, &i)

	table := make([]string, n)
	for j := 0; j < n; j++ {
		fmt.Fscanln(reader, &table[j])
	}

	sort.Slice(table, func(j, k int) bool {
		return table[j] < table[k]
	})
	fmt.Println(table[i-1])
}
