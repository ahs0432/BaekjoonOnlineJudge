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

	var sum int
	var s string
	for {
		s, _ = reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")

		if s == "#" {
			break
		}

		sum = 0
		for i, c := range s {
			if c == ' ' {
				continue
			}

			sum += (i + 1) * int(c-'A'+1)
		}
		fmt.Fprintln(writer, sum)
	}
	writer.Flush()
}
