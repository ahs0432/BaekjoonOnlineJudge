package main

import (
	"bufio"
	"fmt"
	"os"
)

func bubleSort(n, k int, table []int) int {
	count := 0

	var tmp int
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if table[j] > table[j+1] {
				tmp = table[j+1]
				table[j+1] = table[j]
				table[j] = tmp
				count++
			}

			if count == k {
				return j
			}
		}
	}

	return -1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	table := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &table[i])
	}

	count := bubleSort(n, k, table)
	if count == -1 {
		fmt.Fprintln(writer, count)
	} else {
		for i := 0; i < n; i++ {
			fmt.Fprint(writer, table[i], " ")
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
