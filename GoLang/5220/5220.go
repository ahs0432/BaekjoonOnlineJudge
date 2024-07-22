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

	var a, b, cnt int
	for i := 0; i < n; i++ {
		cnt = 0
		fmt.Fscanln(reader, &a, &b)

		for a != 0 {
			if a%2 == 1 {
				cnt++
			}
			a /= 2
		}

		if cnt%2 == b {
			fmt.Fprintln(writer, "Valid")
		} else {
			fmt.Fprintln(writer, "Corrupt")
		}
	}
	writer.Flush()
}
