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

	var s string
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s)
		fmt.Fprintln(writer, len(s))
	}
	writer.Flush()
}
