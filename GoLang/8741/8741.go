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

	for i := 0; i < n; i++ {
		fmt.Fprint(writer, 1)
	}

	for i := 0; i < n-1; i++ {
		fmt.Fprint(writer, 0)
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
