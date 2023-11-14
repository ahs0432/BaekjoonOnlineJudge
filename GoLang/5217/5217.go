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

	var tmp int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)
		fmt.Fprintf(writer, "Pairs for %d: ", tmp)

		s := 1
		for j := 0; j < (tmp-1)/2; j++ {
			if j != 0 {
				fmt.Fprintf(writer, ", ")
			}
			fmt.Fprintf(writer, "%d %d", s, tmp-s)
			s += 1
		}
		fmt.Fprintln(writer)
	}

	writer.Flush()
}
