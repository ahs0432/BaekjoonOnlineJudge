package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var s string
	fmt.Fscanln(reader, &s)
	for i := 0; i < len(s); i++ {
		fmt.Fprintln(writer, i)
	}
	writer.Flush()
}
