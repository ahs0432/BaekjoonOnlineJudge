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

	fmt.Fprint(writer, "AKA")
	for i := 0; i < n; i++ {
		fmt.Fprint(writer, "RAKA")
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
