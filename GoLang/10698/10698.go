package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var x, y, z int
	var o, e string
	var tans int

	var n int
	fmt.Fscanln(reader, &n)

	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &x, &o, &y, &e, &z)

		if o == "+" {
			tans = x + y
		} else {
			tans = x - y
		}

		if tans == z {
			fmt.Fprintf(writer, "Case %d: YES\n", i)
		} else {
			fmt.Fprintf(writer, "Case %d: NO\n", i)
		}
	}
	writer.Flush()
}
