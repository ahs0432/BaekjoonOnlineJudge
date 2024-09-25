package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var prev rune
	var s string
	var max, count int
	for i := 0; i < 3; i++ {
		prev = '-'
		max, count = 0, 0
		fmt.Fscanln(reader, &s)

		for _, c := range s {
			if prev == c {
				count++
			} else {
				prev = c
				count = 1
			}

			if count > max {
				max = count
			}
		}

		fmt.Fprintln(writer, max)
	}
	writer.Flush()
}
