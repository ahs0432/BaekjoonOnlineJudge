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

	var tmp, n, m, x, s, b int
	for i := 0; i < t; i++ {
		tmp = 0
		fmt.Fscanln(reader, &tmp)
		fmt.Fscanln(reader, &n, &m)

		s = 0
		for j := 0; j < n; j++ {
			if j == n-1 {
				fmt.Fscanln(reader, &x)
			} else {
				fmt.Fscan(reader, &x)
			}
			s = max(s, x)
		}

		b = 0
		for j := 0; j < m; j++ {
			if j == m-1 {
				fmt.Fscanln(reader, &x)
			} else {
				fmt.Fscan(reader, &x)
			}
			b = max(b, x)
		}

		if s < b {
			fmt.Fprintln(writer, "B")
		} else {
			fmt.Fprintln(writer, "S")
		}
	}
	writer.Flush()
}
