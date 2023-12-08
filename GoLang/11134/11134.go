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

	var n, c int
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n, &c)

		if n%c == 0 {
			fmt.Fprintln(writer, n/c)
		} else {
			fmt.Fprintln(writer, n/c+1)
		}
	}

	writer.Flush()
}
