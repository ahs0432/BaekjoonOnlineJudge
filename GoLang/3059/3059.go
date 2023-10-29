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
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s)
		t := make([]bool, 26)

		for _, c := range s {
			t[c-65] = true
		}

		sum := 0
		for i, b := range t {
			if !b {
				sum += (65 + i)
			}
		}
		fmt.Fprintln(writer, sum)
	}
	writer.Flush()
}
