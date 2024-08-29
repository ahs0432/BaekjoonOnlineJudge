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

	var n, m int
	var count int
	for i := 0; i < t; i++ {
		count = 0
		fmt.Fscanln(reader, &n, &m)

		for x := 1; x <= n; x++ {
			for y := x + 1; y < n; y++ {
				if ((x*x)+(y*y)+m)%(x*y) == 0 {
					count++
				}
			}
		}

		fmt.Fprintln(writer, count)
	}
	writer.Flush()
}
