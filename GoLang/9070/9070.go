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

	var n int
	var w, c float64
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n)

		ans := 0.0
		bestC := 100001.0
		for j := 0; j < n; j++ {
			fmt.Fscanln(reader, &w, &c)
			if bestC > c/w {
				bestC = c / w
				ans = c
			} else if bestC == c/w && ans > c {
				ans = c
			}
		}
		fmt.Fprintf(writer, "%0.f\n", ans)
	}
	writer.Flush()
}
