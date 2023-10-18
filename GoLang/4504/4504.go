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

	var tmp int
	for {
		fmt.Fscanln(reader, &tmp)

		if tmp == 0 {
			break
		}

		if tmp%n == 0 {
			fmt.Fprintf(writer, "%d is a multiple of %d.\n", tmp, n)
		} else {
			fmt.Fprintf(writer, "%d is NOT a multiple of %d.\n", tmp, n)
		}
	}
	writer.Flush()
}
