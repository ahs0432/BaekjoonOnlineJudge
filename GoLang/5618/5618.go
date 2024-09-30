package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	min := 2147483647
	table := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &table[i])

		if table[i] < min {
			min = table[i]
		}
	}

	count := 0
	for i := 1; i <= min; i++ {
		count = 0
		for _, t := range table {
			if t%i == 0 {
				count++
			} else {
				break
			}
		}

		if count == n {
			fmt.Fprintln(writer, i)
		}
	}
	writer.Flush()
}
