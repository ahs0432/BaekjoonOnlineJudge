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

	var t, a, b int
	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &t)

		for j := 0; j <= t; j++ {
			fmt.Fscanln(reader, &a, &b)
		}
		fmt.Fprintf(writer, "Material Management %d\n", i)
		fmt.Fprintln(writer, "Classification ---- End!")
	}
	writer.Flush()
}
