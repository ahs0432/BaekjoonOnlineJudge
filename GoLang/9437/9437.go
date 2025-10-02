package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, p int
	for {
		fmt.Fscanln(reader, &n, &p)

		if n == 0 {
			break
		}

		for i := 0; i < n/4; i++ {
			if 2*i+1 == p {
				fmt.Fprintln(writer, 2*i+2, n-2*i-1, n-2*i)
			} else if 2*i+2 == p {
				fmt.Fprintln(writer, 2*i+1, n-2*i-1, n-2*i)
			} else if n-2*i-1 == p {
				fmt.Fprintln(writer, 2*i+1, 2*i+2, n-2*i)
			} else if n-2*i == p {
				fmt.Fprintln(writer, 2*i+1, 2*i+2, n-2*i-1)
			}
		}
	}
	writer.Flush()
}
