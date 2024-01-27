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

	s := ""
	for i := 0; i < n; i++ {
		s += "@"
	}
	for i := 0; i < n*4; i++ {
		fmt.Fprintln(writer, s)
	}

	s = ""
	for i := 0; i < n*5; i++ {
		s += "@"
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, s)
	}

	writer.Flush()
}
