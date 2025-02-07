package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t, k int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &k)
		fmt.Fprintln(writer, k*23)
	}
	writer.Flush()
}
