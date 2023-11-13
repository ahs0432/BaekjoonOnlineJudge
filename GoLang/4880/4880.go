package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var a, b, c int
	for {
		fmt.Fscanln(reader, &a, &b, &c)

		if a == 0 && b == 0 && c == 0 {
			break
		} else if b-a == c-b {
			fmt.Fprintln(writer, "AP", c+c-b)
		} else {
			fmt.Fprintln(writer, "GP", c*(c/b))
		}
	}
	writer.Flush()
}
