package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if b < a {
		return a
	} else {
		return b
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var a, b, c int
	for {
		a, b, c = -1, -1, -1
		fmt.Fscanln(reader, &a, &b, &c)

		if a == -1 || b == -1 || c == -1 {
			break
		}

		fmt.Fprintln(writer, max(b-a, c-b)-1)
	}
	writer.Flush()
}
