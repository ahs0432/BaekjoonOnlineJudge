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

	for i := 0; i < n; i++ {
		var p, t int
		fmt.Fscanln(reader, &p, &t)
		fmt.Fprintln(writer, (p - (t / 7) + (t / 4)))
	}
	writer.Flush()
}
