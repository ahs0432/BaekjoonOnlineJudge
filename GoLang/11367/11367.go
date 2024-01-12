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

	var name string
	var s int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &name, &s)

		fmt.Fprint(writer, name)
		if s >= 97 {
			fmt.Fprintln(writer, " A+")
		} else if s >= 90 {
			fmt.Fprintln(writer, " A")
		} else if s >= 87 {
			fmt.Fprintln(writer, " B+")
		} else if s >= 80 {
			fmt.Fprintln(writer, " B")
		} else if s >= 77 {
			fmt.Fprintln(writer, " C+")
		} else if s >= 70 {
			fmt.Fprintln(writer, " C")
		} else if s >= 67 {
			fmt.Fprintln(writer, " D+")
		} else if s >= 60 {
			fmt.Fprintln(writer, " D")
		} else {
			fmt.Fprintln(writer, " F")
		}
	}
	writer.Flush()
}
