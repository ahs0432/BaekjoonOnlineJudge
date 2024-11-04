package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var a, b string
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)

	ta := make([]int, len(a))
	tb := make([]int, len(b))
	for i := 0; i < len(a); i++ {
		ta[i] = int(a[i] - '0')
		tb[i] = int(b[i] - '0')
	}

	for i := 0; i < len(a); i++ {
		fmt.Fprint(writer, ta[i]&tb[i])
	}
	fmt.Fprintln(writer)

	for i := 0; i < len(a); i++ {
		fmt.Fprint(writer, ta[i]|tb[i])
	}
	fmt.Fprintln(writer)

	for i := 0; i < len(a); i++ {
		fmt.Fprint(writer, ta[i]^tb[i])
	}
	fmt.Fprintln(writer)

	for i := 0; i < len(a); i++ {
		if ta[i] == 0 {
			fmt.Fprint(writer, 1)
		} else {
			fmt.Fprint(writer, 0)
		}
	}
	fmt.Fprintln(writer)

	for i := 0; i < len(a); i++ {
		if tb[i] == 0 {
			fmt.Fprint(writer, 1)
		} else {
			fmt.Fprint(writer, 0)
		}
	}
	fmt.Fprintln(writer)

	writer.Flush()
}
