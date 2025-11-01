package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, r, c int
	fmt.Fscanln(reader, &n, &r, &c)

	if (r+c)%2 == 0 {
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				for j := 0; j < n/2; j++ {
					fmt.Fprint(writer, "v.")
				}
				for j := 0; j < n%2; j++ {
					fmt.Fprint(writer, "v")
				}
			} else {
				for j := 0; j < n/2; j++ {
					fmt.Fprint(writer, ".v")
				}
				for j := 0; j < n%2; j++ {
					fmt.Fprint(writer, ".")
				}
			}
			fmt.Fprintln(writer)
		}
	} else {
		for i := 0; i < n; i++ {
			if i%2 == 1 {
				for j := 0; j < n/2; j++ {
					fmt.Fprint(writer, "v.")
				}
				for j := 0; j < n%2; j++ {
					fmt.Fprint(writer, "v")
				}
			} else {
				for j := 0; j < n/2; j++ {
					fmt.Fprint(writer, ".v")
				}
				for j := 0; j < n%2; j++ {
					fmt.Fprint(writer, ".")
				}
			}
			fmt.Fprintln(writer)
		}
	}
	writer.Flush()
}
