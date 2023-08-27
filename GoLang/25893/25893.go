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

	var a, b, c int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b, &c)
		fmt.Fprintln(writer, a, b, c)

		if a < 10 && b < 10 && c < 10 {
			fmt.Fprintln(writer, "zilch")
		} else if a >= 10 && b >= 10 && c >= 10 {
			fmt.Fprintln(writer, "triple-double")
		} else if (a >= 10 && b >= 10) || (a >= 10 && c >= 10) || (b >= 10 && c >= 10) {
			fmt.Fprintln(writer, "double-double")
		} else {
			fmt.Fprintln(writer, "double")
		}
		fmt.Fprintln(writer)
	}

	writer.Flush()
}
