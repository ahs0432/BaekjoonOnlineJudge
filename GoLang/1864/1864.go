package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var s string
	var now, oct, ans int
	for {
		fmt.Fscanln(reader, &s)

		if s == "#" {
			break
		}

		oct = 1
		ans = 0
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == '-' {
				now = 0
			} else if s[i] == '\\' {
				now = 1
			} else if s[i] == '(' {
				now = 2
			} else if s[i] == '@' {
				now = 3
			} else if s[i] == '?' {
				now = 4
			} else if s[i] == '>' {
				now = 5
			} else if s[i] == '&' {
				now = 6
			} else if s[i] == '%' {
				now = 7
			} else if s[i] == '/' {
				now = -1
			}

			ans += (now * oct)
			oct *= 8
		}
		fmt.Fprintln(writer, ans)
	}
	writer.Flush()
}
