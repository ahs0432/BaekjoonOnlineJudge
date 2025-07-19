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

	var a, b, x int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b, &x)
		fmt.Fprintln(writer, a*(x-1)+b)
	}
	writer.Flush()
}
