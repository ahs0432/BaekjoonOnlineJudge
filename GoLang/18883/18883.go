package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	for i := 1; i <= n*m; i++ {
		if i%m == 0 {
			fmt.Fprintln(writer, i)
		} else {
			fmt.Fprintf(writer, "%d ", i)
		}
	}
	writer.Flush()
}
