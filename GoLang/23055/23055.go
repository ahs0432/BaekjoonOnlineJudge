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

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 || i == n {
				fmt.Fprint(writer, "*")
			} else if j == 1 || j == n {
				fmt.Fprint(writer, "*")
			} else if i == j || i+j-1 == n {
				fmt.Fprint(writer, "*")
			} else {
				fmt.Fprint(writer, " ")
			}
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
