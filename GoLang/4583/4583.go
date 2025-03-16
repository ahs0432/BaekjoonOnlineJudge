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
	var ans string

	for {
		fmt.Fscanln(reader, &s)

		if s == "#" {
			break
		}

		ans = ""
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == 'b' {
				ans += "d"
			} else if s[i] == 'd' {
				ans += "b"
			} else if s[i] == 'p' {
				ans += "q"
			} else if s[i] == 'q' {
				ans += "p"
			} else if s[i] == 'i' || s[i] == 'o' || s[i] == 'v' || s[i] == 'w' || s[i] == 'x' {
				ans += string(s[i])
			} else {
				ans = "INVALID"
				break
			}
		}
		fmt.Fprintln(writer, ans)
	}
	writer.Flush()
}
