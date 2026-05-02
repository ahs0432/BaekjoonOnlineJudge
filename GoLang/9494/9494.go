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

	var n, now int
	var s string

	for {
		fmt.Fscanln(reader, &n)
		if n == 0 {
			break
		}

		now = 0
		for i := 0; i < n; i++ {
			s, _ = reader.ReadString('\n')
			s = strings.TrimSuffix(s, "\n")

			for j := now; j < len(s); j++ {
				if s[j] == ' ' {
					break
				}
				now++
			}
		}
		fmt.Fprintln(writer, now+1)
	}
	writer.Flush()
}
