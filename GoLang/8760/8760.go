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
		var w, k int
		fmt.Fscanln(reader, &w, &k)
		fmt.Fprintln(writer, w*k/2)
	}

	writer.Flush()
}
