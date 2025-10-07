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
	fmt.Fprintf(writer, "%d.", int(a/b))

	a = a % b
	for i := 0; a != 0 && i < 1000; i++ {
		a *= 10
		fmt.Fprint(writer, a/b)
		a = a % b
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
