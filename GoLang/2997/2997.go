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

	a := table[1] - table[0]
	b := table[2] - table[1]

	if a == b {
		fmt.Println(table[2] + a)
	} else if a < b {
		fmt.Println(table[1] + a)
	} else {
		fmt.Println(table[0] + b)
	}
}
