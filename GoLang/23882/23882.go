package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	table := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &table[i])
	}

	check := false
	var count, max, tmp int
	for i := n - 1; i > 0; i-- {
		max = 0

		for j := 0; j <= i; j++ {
			if table[j] > table[max] {
				max = j
			}
		}

		if max != i {
			tmp = table[i]
			table[i] = table[max]
			table[max] = tmp

			count++
		}

		if count == k {
			check = true
			for j := 0; j < n; j++ {
				fmt.Fprintf(writer, "%d ", table[j])
			}
		}
	}

	if !check {
		fmt.Fprint(writer, -1)
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
