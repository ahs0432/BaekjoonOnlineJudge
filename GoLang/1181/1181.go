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

	set := map[string]bool{}

	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscanln(reader, &tmp)
		set[tmp] = true
	}

	table := make([]string, 0)
	for k := range set {
		table = append(table, k)
	}

	sort.Slice(table, func(i, j int) bool {
		if len(table[i]) == len(table[j]) {
			return table[i] < table[j]
		}
		return len(table[i]) < len(table[j])
	})

	for _, t := range table {
		fmt.Println(t)
	}
}
