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
		var s string
		fmt.Fscanln(reader, &s)

		now := byte('-')
		for j := 0; j < len(s); j++ {
			if s[j] == now {
				continue
			}
			now = s[j]
			fmt.Fprint(writer, string(s[j]))
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
