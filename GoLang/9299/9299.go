package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var n int
	var table []int
	for i := 1; i <= t; i++ {
		fmt.Fscan(reader, &n)
		table = make([]int, n+1)
		for j := n; j >= 0; j-- {
			fmt.Fscan(reader, &table[j])
		}

		fmt.Fprintf(writer, "Case %d: %d", i, n-1)
		for j := n; j > 0; j-- {
			fmt.Fprintf(writer, " %d", table[j]*j)
		}
		fmt.Fprintln(writer)
	}

	writer.Flush()
}
