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

	for i := 0; i < n+2; i++ {
		fmt.Fprint(writer, "@")
	}
	fmt.Fprintln(writer)

	s := "@"
	for i := 0; i < n; i++ {
		s += " "
	}
	s += "@"

	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, s)
	}

	for i := 0; i < n+2; i++ {
		fmt.Fprint(writer, "@")
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
