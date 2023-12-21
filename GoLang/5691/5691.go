package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var a, b int

	for {
		fmt.Fscanln(reader, &a, &b)

		if a == 0 && b == 0 {
			break
		}
		fmt.Fprintln(writer, 2*a-b)
	}
	writer.Flush()
}
