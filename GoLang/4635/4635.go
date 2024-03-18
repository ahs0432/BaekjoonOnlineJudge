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
	var m, t int
	var p, total int
	for {
		p, total = 0, 0
		fmt.Fscanln(reader, &n)
		if n == -1 {
			break
		}

		for i := 0; i < n; i++ {
			fmt.Fscanln(reader, &m, &t)
			total += (m * (t - p))
			p = t
		}
		fmt.Fprintln(writer, total, "miles")
	}

	writer.Flush()
}
