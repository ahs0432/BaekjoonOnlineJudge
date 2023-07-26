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
		var a, b string
		fmt.Fscanln(reader, &a, &b)

		if a == b {
			fmt.Fprintln(writer, "OK")
		} else {
			fmt.Fprintln(writer, "ERROR")
		}
	}
	writer.Flush()
}
