package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	table := make([]float64, 6)
	var sum float64
	for {
		sum = 0
		fmt.Fscanln(reader, &table[0], &table[1], &table[2], &table[3], &table[4], &table[5])

		if table[0] == 0 && table[1] == 0 && table[2] == 0 && table[3] == 0 && table[4] == 0 && table[5] == 0 {
			break
		}

		sort.Slice(table, func(i, j int) bool {
			return table[i] < table[j]
		})

		for i := 1; i < len(table)-1; i++ {
			sum += table[i]
		}

		fmt.Fprintln(writer, sum/4)
	}
	writer.Flush()
}
