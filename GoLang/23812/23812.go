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

	s1 := ""
	for i := 0; i < n; i++ {
		s1 += "@"
	}
	for i := 0; i < 3*n; i++ {
		s1 += " "
	}
	for i := 0; i < n; i++ {
		s1 += "@"
	}

	s2 := ""
	for i := 0; i < n*5; i++ {
		s2 += "@"
	}

	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, s1)
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, s2)
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, s1)
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, s2)
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, s1)
	}

	writer.Flush()
}
