package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, sum int
	fmt.Fscanln(reader, &n)

	var s string
	for i := 0; i < n; i++ {
		sum = 0
		s, _ = reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")

		for _, c := range s {
			if c >= 'A' && c <= 'Z' {
				sum += int(c - 'A' + 1)
			}
		}

		if sum == 100 {
			fmt.Fprintln(writer, "PERFECT LIFE")
		} else {
			fmt.Fprintln(writer, sum)
		}
	}
	writer.Flush()
}
