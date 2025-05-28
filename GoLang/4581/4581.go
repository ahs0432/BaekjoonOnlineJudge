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
	var y, n, a, l int
	for {
		fmt.Fscanln(reader, &s)

		if s == "#" {
			break
		}

		y, n, a, l = 0, 0, 0, len(s)
		for _, c := range s {
			if c == 'Y' {
				y++
			} else if c == 'N' {
				n++
			} else if c == 'A' {
				a++
			}
		}

		if l%2 == 1 {
			l++
		}

		if a >= l/2 {
			fmt.Fprintln(writer, "need quorum")
		} else if y == n {
			fmt.Fprintln(writer, "tie")
		} else if y > n {
			fmt.Fprintln(writer, "yes")
		} else if y < n {
			fmt.Fprintln(writer, "no")
		}
	}
	writer.Flush()
}
