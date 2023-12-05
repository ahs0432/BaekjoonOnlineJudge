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

	table := make([]int, 0)

	var tmp int
	fmt.Fscanln(reader, &tmp)
	table = append(table, tmp)

	var a, b int
	t := true
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b)

		if table[i]+a-b < 0 {
			t = false
		}

		table = append(table, table[i]+a-b)
	}

	if t {
		sort.Slice(table, func(i, j int) bool {
			return table[i] > table[j]
		})
		fmt.Println(table[0])
	} else {
		fmt.Println(0)
	}
}
