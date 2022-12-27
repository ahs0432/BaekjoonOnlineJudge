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
		var n, m int
		fmt.Fscanln(reader, &n, &m)
		if n == 0 && m == 0 {
			break
		}

		if n > m {
			fmt.Fprintln(writer, "Yes")
		} else {
			fmt.Fprintln(writer, "No")
		}
	}
	writer.Flush()
}
