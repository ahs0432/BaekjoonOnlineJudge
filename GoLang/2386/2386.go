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
		sum = 0
		s, _ = reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")

		if s == "#" {
			break
		}

		for i := 2; i < len(s); i++ {
			if s[0] == s[i] || s[0]-32 == s[i] {
				sum++
			}
		}
		fmt.Fprintln(writer, string(s[0]), sum)
	}
	writer.Flush()
}
