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

	var s string
	var checkNum bool
	var upper, lower int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s)
		if len(s) > 10 {
			continue
		}

		checkNum = true
		upper, lower = 0, 0
		for _, c := range s {
			if c >= '0' && c <= '9' {
				continue
			}

			if 'A' <= c && c <= 'Z' {
				upper++
			} else if 'a' <= c && c <= 'z' {
				lower++
			}
			checkNum = false
		}

		if upper > lower {
			continue
		}

		if checkNum {
			continue
		}
		fmt.Fprintln(writer, s)
	}
	writer.Flush()
}
