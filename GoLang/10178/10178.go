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

	var a, b int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b)
		fmt.Fprintf(writer, "You get %d piece(s) and your dad gets %d piece(s).\n", a/b, a%b)
	}
	writer.Flush()
}
