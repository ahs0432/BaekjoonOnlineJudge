package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, q int
	fmt.Fscanln(reader, &n, &q)

	table := make([]int, n+1)
	var t, p, x int
	for i := 0; i < q; i++ {
		fmt.Fscanln(reader, &t, &p, &x)

		if t == 1 {
			for j := p; j <= n; j++ {
				table[j] += x
			}
		} else {
			fmt.Fprintln(writer, table[x]-table[p-1])
		}
	}
	writer.Flush()
}
