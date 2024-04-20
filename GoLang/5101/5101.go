package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var a, b, c, ans int
	for {
		fmt.Fscanln(reader, &a, &b, &c)
		if a == 0 && b == 0 && c == 0 {
			break
		}

		ans = c - a + b
		if (c-a)%b == 0 && (c-a)/b >= 0 {
			fmt.Fprintln(writer, ans/b)
		} else {
			fmt.Fprintln(writer, "X")
		}
	}
	writer.Flush()
}
