package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, tmp int
	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)

		fmt.Fprintf(writer, "%d %d %d %d\n", tmp/25, tmp%25/10, tmp%25%10/5, tmp%25%10%5)
	}
	writer.Flush()
}
