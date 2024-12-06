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

	var n int
	fmt.Fscanln(reader, &n)

	var s string
	var countG, countB int
	for i := 0; i < n; i++ {
		s, _ = reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")

		countG, countB = 0, 0

		for _, c := range s {
			if c == 'G' || c == 'g' {
				countG++
			} else if c == 'B' || c == 'b' {
				countB++
			}
		}

		if countB == countG {
			fmt.Fprintln(writer, s, "is NEUTRAL")
		} else if countB < countG {
			fmt.Fprintln(writer, s, "is GOOD")
		} else {
			fmt.Fprintln(writer, s, "is A BADDY")
		}
	}

	writer.Flush()
}
