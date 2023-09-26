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

	var n, cnt int
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n)

		cnt = 0
		for n != 0 {
			if n%2 == 1 {
				fmt.Fprintf(writer, "%d ", cnt)
			}
			n /= 2
			cnt++
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
