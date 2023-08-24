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
		var s string
		var a, b int

		fmt.Fscanln(reader, &s, &a, &b)
		s = s[0:a] + s[b:]
		fmt.Fprintln(writer, s)
	}

	writer.Flush()
}
