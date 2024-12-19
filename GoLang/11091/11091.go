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
	var checks []bool
	var ans string
	for i := 0; i < n; i++ {
		s, _ = reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")

		checks = make([]bool, 26)
		for _, c := range s {
			if c >= 'A' && c <= 'Z' {
				checks[int(c-'A')] = true
			} else if c >= 'a' && c <= 'z' {
				checks[int(c-'a')] = true
			}
		}

		ans = ""
		for j, check := range checks {
			if !check {
				ans += string(rune('a' + j))
			}
		}

		if ans == "" {
			fmt.Fprintln(writer, "pangram")
		} else {
			fmt.Fprintln(writer, "missing", ans)
		}
	}
	writer.Flush()
}
