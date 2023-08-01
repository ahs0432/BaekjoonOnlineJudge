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
		var tmp string
		fmt.Fscanln(reader, &tmp)

		s := 0
		for _, c := range tmp {
			if c == 'U' {
				s++
			} else if c == 'D' {
				break
			}
		}

		fmt.Fprintln(writer, s)
	}

	writer.Flush()
}
