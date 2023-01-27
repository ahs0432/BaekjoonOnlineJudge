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
		var m int
		fmt.Fscanln(reader, &m)

		for j := 0; j < m; j++ {
			for k := 0; k < m; k++ {
				if j == 0 || k == 0 || j == m-1 || k == m-1 {
					fmt.Fprint(writer, "#")
				} else {
					fmt.Fprint(writer, "J")
				}
			}
			fmt.Fprintln(writer)
		}
		fmt.Fprintln(writer)
	}

	writer.Flush()
}
