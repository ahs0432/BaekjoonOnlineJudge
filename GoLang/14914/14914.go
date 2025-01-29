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
	fmt.Fscanln(reader, &a, &b)

	for i := 1; a/i != 0 && b/i != 0; i++ {
		if a%i == 0 && b%i == 0 {
			fmt.Fprintln(writer, i, a/i, b/i)
		}
	}
	writer.Flush()
}
