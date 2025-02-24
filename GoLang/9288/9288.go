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
	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &tmp)
		fmt.Fprintf(writer, "Case %d:\n", i)

		for a := 1; a <= 6; a++ {
			for b := a; b <= 6; b++ {
				if a+b == tmp {
					fmt.Fprintf(writer, "(%d,%d)\n", a, b)
				}
			}
		}
	}
	writer.Flush()
}
