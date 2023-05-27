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
	fmt.Fprintln(writer, "Gnomes:")

	for i := 0; i < n; i++ {
		t := make([]int, 3)
		fmt.Fscanln(reader, &t[0], &t[1], &t[2])

		if t[0] < t[1] && t[1] < t[2] {
			fmt.Fprintln(writer, "Ordered")
		} else if t[0] > t[1] && t[1] > t[2] {
			fmt.Fprintln(writer, "Ordered")
		} else {
			fmt.Fprintln(writer, "Unordered")
		}
	}

	writer.Flush()
}
