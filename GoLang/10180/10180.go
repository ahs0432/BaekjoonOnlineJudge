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

	var n, d, v, f, c, count int
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n, &d)

		count = 0
		for j := 0; j < n; j++ {
			fmt.Fscanln(reader, &v, &f, &c)
			if v*f/c >= d {
				count += 1
			}
		}
		fmt.Fprintln(writer, count)
	}
	writer.Flush()
}
