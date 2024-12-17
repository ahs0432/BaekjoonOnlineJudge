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

	var a, b int
	f := make([]int, n)
	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &a, &b)

		f[a-1]++
		f[b-1]++
	}

	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, f[i])
	}
	writer.Flush()
}
