package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscanln(reader, &s)

	answer := 0
	for i := 0; i < len(s); i++ {
		answer++

		if i >= len(s)-1 {
			continue
		}

		if (string(s[i]) == "c" && (string(s[i+1]) == "=" || string(s[i+1]) == "-")) ||
			((string(s[i]) == "l" || string(s[i]) == "n") && string(s[i+1]) == "j") ||
			((string(s[i]) == "s" || string(s[i]) == "z") && string(s[i+1]) == "=") {
			i++
		} else if string(s[i]) == "d" {
			if string(s[i+1]) == "-" {
				i++
			} else if i+2 < len(s) && string(s[i+1]) == "z" && string(s[i+2]) == "=" {
				i += 2
			}
		}
	}

	fmt.Println(answer)
}
