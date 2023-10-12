package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t, n, c, total_c int
	var g, total_g float64
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n)
		total_c, total_g = 0, 0.0
		for j := 0; j < n; j++ {
			fmt.Fscanln(reader, &c, &g)
			total_c += c
			total_g += (float64(c) * g)
		}
		fmt.Fprintf(writer, "%d %.1f\n", total_c, total_g/float64(total_c))
	}
	writer.Flush()
}
