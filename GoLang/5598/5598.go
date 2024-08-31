package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var s string
	fmt.Fscanln(reader, &s)

	for _, c := range s {
		if c >= 'A' && c <= 'C' {
			fmt.Fprint(writer, string(c+23))
		} else {
			fmt.Fprint(writer, string(c-3))
		}
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
