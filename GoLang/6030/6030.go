package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var p, q int
	fmt.Fscanln(reader, &p, &q)

	var ps, qs []int
	for i := 1; i < max(p, q)+1; i++ {
		if p%i == 0 {
			ps = append(ps, i)
		}

		if q%i == 0 {
			qs = append(qs, i)
		}
	}

	for _, pi := range ps {
		for _, qi := range qs {
			fmt.Fprintln(writer, pi, qi)
		}
	}
	writer.Flush()
}
