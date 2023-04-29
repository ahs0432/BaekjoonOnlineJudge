package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for {
		var n, m int
		fmt.Fscanln(reader, &n, &m)

		if n == 0 && m == 0 {
			break
		}
		fmt.Fprintln(writer, n+m)
	}
	writer.Flush()
}
