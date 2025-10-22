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

	var h, w int
	var s string
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &h, &w)

		for j := 0; j < h; j++ {
			fmt.Fscanln(reader, &s)

			for k := len(s) - 1; k >= 0; k-- {
				fmt.Fprint(writer, string(s[k]))
			}
			fmt.Fprintln(writer)
		}
	}
	writer.Flush()
}
