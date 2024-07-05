package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var a, b, n int
	fmt.Fscanln(reader, &a, &b, &n)

	for i := 1; i <= n; i++ {
		fmt.Fprint(writer, a*n+b*i, " ")
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
