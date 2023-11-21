package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n float64
	prev := -11.0
	for {
		fmt.Fscanln(reader, &n)

		if n == 999 {
			break
		}

		if prev != -11 {
			fmt.Fprintf(writer, "%.2f\n", n-prev)
		}
		prev = n
	}
	writer.Flush()
}
