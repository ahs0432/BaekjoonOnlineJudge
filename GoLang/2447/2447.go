package main

import (
	"bufio"
	"fmt"
	"os"
)

func draw(writer *bufio.Writer, n, x, y int) {
	if n == 1 {
		fmt.Fprint(writer, "*")
	} else if (x/(n/3))%3 == 1 && (y/(n/3))%3 == 1 {
		fmt.Fprintf(writer, " ")
	} else {
		draw(writer, n/3, x, y)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			draw(writer, n, i, j)
		}
		fmt.Fprintln(writer)
	}

	writer.Flush()
}
