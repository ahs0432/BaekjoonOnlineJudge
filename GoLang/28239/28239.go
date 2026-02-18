package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var m uint64
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &m)

		onesCount := bits.OnesCount64(m)

		if onesCount == 1 {
			j := bits.TrailingZeros64(m)
			if j > 0 {
				fmt.Fprintf(writer, "%d %d", j-1, j-1)
			}
		} else {
			first := true
			for j := 0; j < 60; j++ {
				if (m & (1 << uint(j))) != 0 {
					if !first {
						fmt.Fprint(writer, " ")
					}
					fmt.Fprintf(writer, "%d", j)
					first = false
				}
			}
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
