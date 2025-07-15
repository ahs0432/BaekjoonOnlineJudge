package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var l, r, s int
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &l, &r, &s)

		if s-l < r-s {
			fmt.Fprintln(writer, (s-l)*2+1)
		} else {
			fmt.Fprintln(writer, (r-s)*2)
		}
	}
	writer.Flush()
}
