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

	s := "int"
	for n != 0 {
		n -= 4
		s = "long " + s
	}
	fmt.Fprintln(writer, s)
	writer.Flush()
}
