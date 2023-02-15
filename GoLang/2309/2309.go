package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	total := 0
	table := make([]int, 9)
	for i := 0; i < 9; i++ {
		fmt.Fscanln(reader, &table[i])
		total += table[i]
	}

	sort.Slice(table, func(i, j int) bool {
		return table[i] < table[j]
	})

	for i := 0; i < 9; i++ {
		for j := i; j < 9; j++ {
			if (total - table[i] - table[j]) == 100 {
				for k, t := range table {
					if k == i || k == j {
						continue
					}
					fmt.Println(t)
				}

				return
			}
		}
	}
}
