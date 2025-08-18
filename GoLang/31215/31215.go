package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t, n int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n)

		if n >= 3 {
			fmt.Fprintln(writer, 3)
		} else {
			fmt.Fprintln(writer, 1)
		}
	}
	writer.Flush()
}
