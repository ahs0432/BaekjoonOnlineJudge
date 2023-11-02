package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var n, m int
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n, &m)
		fmt.Fprintln(writer, (m*2)-n, n-m)
	}
	writer.Flush()
}
