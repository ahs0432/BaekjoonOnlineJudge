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

	basket := make([]int, n)

	var i, j, k int
	for a := 0; a < m; a++ {
		fmt.Fscanln(reader, &i, &j, &k)

		for i = i - 1; i < j; i++ {
			basket[i] = k
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprint(writer, basket[i], " ")
		if i == n-1 {
			fmt.Fprintln(writer)
		}
	}

	writer.Flush()
}
