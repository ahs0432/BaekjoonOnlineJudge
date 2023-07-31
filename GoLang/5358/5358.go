package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for {
		s, _, _ := reader.ReadLine()
		if len(s) == 0 {
			break
		}

		for i := 0; i < len(s); i++ {
			if s[i] == 'i' {
				s[i] = 'e'
			} else if s[i] == 'I' {
				s[i] = 'E'
			} else if s[i] == 'e' {
				s[i] = 'i'
			} else if s[i] == 'E' {
				s[i] = 'I'
			}
		}

		fmt.Fprintln(writer, string(s))
	}
	writer.Flush()
}
