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

	a := map[int]bool{}
	for i := 0; i < n; i++ {
		var tmp int
		if i == n-1 {
			fmt.Fscanln(reader, &tmp)
		} else {
			fmt.Fscan(reader, &tmp)
		}

		a[tmp] = true
	}

	var m int
	fmt.Fscanln(reader, &m)
	for i := 0; i < m; i++ {
		var tmp int
		if i == m-1 {
			fmt.Fscanln(reader, &tmp)
		} else {
			fmt.Fscan(reader, &tmp)
		}

		if _, e := a[tmp]; e {
			fmt.Fprintln(writer, 1)
		} else {
			fmt.Fprintln(writer, 0)
		}
	}
	writer.Flush()
}
