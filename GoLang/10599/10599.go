package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var a, b, c, d int
	for {
		fmt.Fscanln(reader, &a, &b, &c, &d)

		if a == 0 && b == 0 && c == 0 && d == 0 {
			break
		}

		fmt.Fprintln(writer, c-b, d-a)
	}
	writer.Flush()
}
