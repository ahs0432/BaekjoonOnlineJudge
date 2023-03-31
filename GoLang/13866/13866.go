package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	table := make([]int, 4)
	fmt.Fscanln(reader, &table[0], &table[1], &table[2], &table[3])

	sort.Slice(table, func(i, j int) bool {
		return table[i] < table[j]
	})

	if table[0]+table[3] < table[1]+table[2] {
		fmt.Println((table[1] + table[2]) - (table[0] + table[3]))
	} else {
		fmt.Println((table[0] + table[3]) - (table[1] + table[2]))
	}
}
