package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var a, b, n int
	fmt.Fscanln(reader, &a, &b)
	fmt.Fscanln(reader, &n)

	var e int
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &e)

		tmp = 0
		if e >= 1000 {
			tmp = 1000*a + (e-1000)*b
		} else {
			tmp = e * a
		}

		fmt.Fprintln(writer, e, tmp)
	}
	writer.Flush()
}
