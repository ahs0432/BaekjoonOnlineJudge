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

	p := make([]int, n+m)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i])
	}

	var tmp int
	for i := 0; i < n; i++ {
		for j := 0; j < n+m; j++ {
			fmt.Fscan(reader, &tmp)
			p[i] -= tmp
			p[j] += tmp
		}
	}

	for i := 0; i < n+m; i++ {
		fmt.Fprint(writer, p[i], " ")
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
