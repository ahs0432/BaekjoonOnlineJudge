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
		var m int
		fmt.Fscanln(reader, &m)
		fmt.Fprintln(writer, m*m)
	}

	writer.Flush()
}
