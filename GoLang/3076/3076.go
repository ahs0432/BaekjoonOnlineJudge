package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var r, c, a, b int
	fmt.Fscanln(reader, &r, &c)
	fmt.Fscanln(reader, &a, &b)

	for i := 0; i < r; i++ {
		for j := 0; j < a; j++ {
			for k := 0; k < c; k++ {
				for l := 0; l < b; l++ {
					if (k+i)%2 == 0 {
						fmt.Fprint(writer, "X")
					} else {
						fmt.Fprint(writer, ".")
					}
				}
			}
			fmt.Fprintln(writer)
		}
	}
	writer.Flush()
}
