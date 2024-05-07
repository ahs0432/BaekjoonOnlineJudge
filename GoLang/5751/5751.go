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
	var tmp, a, b int
	for {
		n = 0
		fmt.Fscanln(reader, &n)

		if n == 0 {
			break
		}

		a, b = 0, 0
		for i := 0; i < n; i++ {
			if i == n-1 {
				fmt.Fscanln(reader, &tmp)
			} else {
				fmt.Fscan(reader, &tmp)
			}

			if tmp == 0 {
				a++
			} else {
				b++
			}
		}

		fmt.Fprintf(writer, "Mary won %d times and John won %d times\n", a, b)
	}

	writer.Flush()
}
