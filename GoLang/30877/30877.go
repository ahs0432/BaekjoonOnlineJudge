package main

import (
	"bufio"
	"fmt"
	"os"
)

func rep(c byte) byte {
	if 'a' <= c && c <= 'z' {
		return c - 'a' + 'A'
	}
	return c
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var s1, s2 string
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s1, &s2)

		for j := 0; j < len(s1); j++ {
			if s1[j] == 'x' || s1[j] == 'X' {
				fmt.Fprint(writer, string(rep(s2[j])))
				break
			}
		}
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
