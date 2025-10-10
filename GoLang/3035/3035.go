package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var r, c, zr, zc int
	fmt.Fscanln(reader, &r, &c, &zr, &zc)

	s := make([]string, 51)
	for i := 0; i < r; i++ {
		fmt.Fscanln(reader, &s[i])
	}

	for i := 0; i < r; i++ {
		for j := 0; j < zr; j++ {
			for k := 0; k < c; k++ {
				for l := 0; l < zc; l++ {
					fmt.Fprint(writer, string(s[i][k]))
				}
			}
			fmt.Fprintln(writer)
		}
	}

	writer.Flush()
}
