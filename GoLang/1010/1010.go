package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var k, n, m, dam int
	fmt.Fscanln(reader, &k)

	for i := 0; i < k; i++ {
		dam = 1
		fmt.Fscanln(reader, &n, &m)

		for j := 0; j < n; j++ {
			dam *= m - j
			dam /= 1 + j
		}
		fmt.Fprintln(writer, dam)
	}
	writer.Flush()
}
