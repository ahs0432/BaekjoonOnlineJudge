package main

import (
	"bufio"
	"fmt"
	"os"
)

func calc(x1, y1, x2, y2 int) int {
	return ((x2 - x1) + (y2 - y1)) * 2
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var a, b, c, d int
	fmt.Fscanln(reader, &a, &b, &c, &d)
	fmt.Fprintln(writer, calc(a, b, c, d))

	var tmpA, tmpB, tmpC, tmpD int
	for i := 0; i < n-1; i++ {
		fmt.Fscanln(reader, &tmpA, &tmpB, &tmpC, &tmpD)
		a, b, c, d = min(a, tmpA), min(b, tmpB), max(c, tmpC), max(d, tmpD)
		fmt.Fprintln(writer, calc(a, b, c, d))
	}
	writer.Flush()
}
