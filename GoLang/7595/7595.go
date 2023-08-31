package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for {
		var n int
		fmt.Fscanln(reader, &n)

		if n == 0 {
			break
		}

		s := ""
		for i := 0; i < n; i++ {
			s += "*"
			fmt.Fprintln(writer, s)
		}
	}

	writer.Flush()
}
