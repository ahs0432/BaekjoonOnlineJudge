package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var n int
	var s string
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n, &s)
		fmt.Fprintln(writer, s[:n-1]+s[n:])
	}
	writer.Flush()
}
