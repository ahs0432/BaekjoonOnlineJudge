package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var c, d int
	for {
		fmt.Fscanln(reader, &c, &d)

		if c == 0 && d == 0 {
			break
		}

		fmt.Fprintln(writer, min(min((30*c+40*d), (35*c+30*d)), 40*c+20*d))
	}
	writer.Flush()
}
