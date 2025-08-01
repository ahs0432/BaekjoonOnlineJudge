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

	for i := 1; i <= n; i++ {
		fmt.Fprintf(writer, "Hello World, Judge %d!\n", i)
	}
	writer.Flush()
}
