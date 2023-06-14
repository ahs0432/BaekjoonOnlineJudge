package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	table := make([]int, 3)
	fmt.Fscanln(reader, &table[0], &table[1], &table[2])

	sort.Slice(table, func(i, j int) bool {
		return table[i] < table[j]
	})

	sum := table[2] - table[0]
	sum += (table[2] - table[1])

	fmt.Println(sum)
}
