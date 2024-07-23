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
		for j := 0; j < n; j++ {
			fmt.Fprint(writer, "G")
		}
		for j := 0; j < n*3; j++ {
			fmt.Fprint(writer, ".")
		}
		fmt.Fprintln(writer)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fprint(writer, ".")
		}
		for j := 0; j < n; j++ {
			fmt.Fprint(writer, "I")
		}
		for j := 0; j < n; j++ {
			fmt.Fprint(writer, ".")
		}
		for j := 0; j < n; j++ {
			fmt.Fprint(writer, "T")
		}
		fmt.Fprintln(writer)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n*2; j++ {
			fmt.Fprint(writer, ".")
		}
		for j := 0; j < n; j++ {
			fmt.Fprint(writer, "S")
		}
		for j := 0; j < n; j++ {
			fmt.Fprint(writer, ".")
		}
		fmt.Fprintln(writer)
	}

	writer.Flush()
}
