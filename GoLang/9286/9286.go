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

	var m, tmp int
	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &m)
		fmt.Fprintf(writer, "Case %d:\n", i)
		for j := 0; j < m; j++ {
			fmt.Fscanln(reader, &tmp)
			if tmp < 6 {
				fmt.Fprintln(writer, tmp+1)
			}
		}
	}

	writer.Flush()
}
