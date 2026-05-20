package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var l, w, a int

	for {
		fmt.Fscanln(reader, &l, &w, &a)

		if l == 0 && w == 0 && a == 0 {
			break
		}

		if l == 0 {
			fmt.Fprintln(writer, a/w, w, a)
		} else if w == 0 {
			fmt.Fprintln(writer, l, a/l, a)
		} else {
			fmt.Fprintln(writer, l, w, l*w)
		}
	}
	writer.Flush()
}
