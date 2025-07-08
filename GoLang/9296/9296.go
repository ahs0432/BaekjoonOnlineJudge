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

	var l, count int
	var tmpA, tmpB string
	for i := 0; i < n; i++ {
		count = 0
		fmt.Fscanln(reader, &l)
		fmt.Fscanln(reader, &tmpA)
		fmt.Fscanln(reader, &tmpB)

		for j := 0; j < l; j++ {
			if tmpA[j] != tmpB[j] {
				count++
			}
		}

		fmt.Fprintf(writer, "Case %d: %d\n", i+1, count)
	}
	writer.Flush()
}
