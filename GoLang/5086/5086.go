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
		} else if b%a == 0 {
			fmt.Fprintln(writer, "factor")
		} else if a%b == 0 {
			fmt.Fprintln(writer, "multiple")
		} else {
			fmt.Fprintln(writer, "neither")
		}
	}
	writer.Flush()
}
