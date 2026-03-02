package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, m, l, r int
	for {
		fmt.Fscanln(reader, &n)
		if n == 0 {
			break
		}

		m = 25
		if n == m {
			fmt.Fprintln(writer, m)
		} else {
			l, r = 1, 50

			for n != m {
				fmt.Fprintf(writer, "%d ", m)
				if m > n {
					r = m - 1
				} else {
					l = m + 1
				}
				m = (l + r) / 2
			}
			fmt.Fprintln(writer, m)
		}
	}
	writer.Flush()
}
