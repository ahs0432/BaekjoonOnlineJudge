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
		if c != 'C' && c != 'A' && c != 'M' && c != 'B' && c != 'R' && c != 'I' && c != 'D' && c != 'G' && c != 'E' {
			fmt.Fprint(writer, string(c))
		}
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
