package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	table := make([]int, n)
	tmpTable := make([]int, n)

	for i := 0; i < n; i++ {
		table[i] = i + 1
	}

	var i, j, k, tmpI int
	for a := 0; a < m; a++ {
		fmt.Fscanln(reader, &i, &j, &k)

		i, j, k = i-1, j-1, k-1

		tmpI = i
		for b := 0; b < j-i+1; b++ {
			if k+b <= j {
				tmpTable[b+i] = table[k+b]
			} else {
				tmpTable[b+i] = table[tmpI]
				tmpI++
			}
		}

		for d := i; d <= j; d++ {
			table[d] = tmpTable[d]
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprint(writer, table[i], " ")
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
