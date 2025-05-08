package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var ans int = 0

func checkTable(n, num int, table []int) {
	if num > n {
		return
	}

	ans = max(ans, num)

	num *= 10
	for i := 0; i < len(table); i++ {
		checkTable(n, num+table[i], table)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	table := make([]int, k)
	for i := range table {
		fmt.Fscan(reader, &table[i])
	}

	sort.Slice(table, func(i, j int) bool {
		return table[i] > table[j]
	})
	checkTable(n, 0, table)
	fmt.Println(ans)
}
