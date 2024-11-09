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

	var s string
	var alpha []bool
	var check bool
	for {
		s, _ = reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")

		if s == "*" {
			break
		}

		check = true
		alpha = make([]bool, 26)
		for _, c := range s {
			if c >= 'a' && c <= 'z' {
				alpha[int(c-97)] = true
			}
		}

		for i := 0; i < 26; i++ {
			if !alpha[i] {
				check = false
			}
		}

		if check {
			fmt.Fprintln(writer, "Y")
		} else {
			fmt.Fprintln(writer, "N")
		}
	}
	writer.Flush()
}
