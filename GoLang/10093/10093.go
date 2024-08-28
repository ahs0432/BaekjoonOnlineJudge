package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var a, b int64
	fmt.Fscanln(reader, &a, &b)

	if a > b {
		a, b = b, a
	}

	if a == b {
		fmt.Fprintln(writer, 0)
	} else {
		fmt.Fprintln(writer, b-a-1)
	}
	for a += 1; a < b; a++ {
		fmt.Fprint(writer, a, " ")
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
