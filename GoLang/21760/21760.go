package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var n, m, k, d, b int64
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n, &m, &k, &d)
		b = 0

		for n*m*b*(k*(m-1)+m*(n-1)) <= 2*d {
			b++
		}

		if b <= 1 {
			fmt.Fprintln(writer, -1)
			continue
		}

		fmt.Fprintln(writer, (n*m*(b-1)*(k*(m-1)+m*(n-1)))/2)
	}
	writer.Flush()
}
