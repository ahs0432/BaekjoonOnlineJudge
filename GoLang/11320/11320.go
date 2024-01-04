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

	var a, b float64
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b)
		fmt.Fprintf(writer, "%.0f\n", (a*a/2)/(b*b/2))
	}
	writer.Flush()
}
