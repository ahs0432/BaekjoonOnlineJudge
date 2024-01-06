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
	fmt.Fprintln(writer, "int a;")

	for i := 1; i <= n; i++ {
		fmt.Fprint(writer, "int ")
		for j := 0; j < i; j++ {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprint(writer, "ptr")

		if i != 1 {
			fmt.Fprint(writer, i)
		}
		fmt.Fprint(writer, " = ")

		if i == 1 {
			fmt.Fprintln(writer, "&a;")
		} else {
			fmt.Fprint(writer, "&ptr")
			if i == 2 {
				fmt.Fprintln(writer, ";")
			} else {
				fmt.Fprint(writer, i-1)
				fmt.Fprintln(writer, ";")
			}
		}
	}

	writer.Flush()
}
