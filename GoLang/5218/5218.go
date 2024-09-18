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

	var a, b string
	var dist int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b)

		fmt.Fprint(writer, "Distances: ")
		for j := 0; j < len(a); j++ {
			dist = int(b[j]) - int(a[j])
			if dist < 0 {
				dist += 26
			}
			fmt.Fprint(writer, dist, " ")
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
