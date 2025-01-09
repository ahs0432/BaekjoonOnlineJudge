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
		if i == n && n%2 == 1 {
			fmt.Fprint(writer, "3")
		} else if i%2 == 1 {
			fmt.Fprint(writer, "1 ")
		} else if i%2 == 0 {
			fmt.Fprint(writer, "2 ")
		}

	}
	fmt.Fprintln(writer)
	writer.Flush()
}
