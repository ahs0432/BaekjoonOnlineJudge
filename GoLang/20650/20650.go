package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	table := make([]int, 7)
	fmt.Fscanln(reader, &table[0], &table[1], &table[2], &table[3], &table[4], &table[5], &table[6])

	sort.Slice(table, func(i, j int) bool {
		return table[i] < table[j]
	})

	fmt.Println(table[0], table[1], table[6]-table[0]-table[1])
}
