package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var l, u, n, sp int

loop:
	for {
		l, u, n, sp = 0, 0, 0, 0
		s, err := reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")
		switch err {
		case nil:
		case io.EOF:
			break loop
		}
		for _, b := range s {
			if b >= 'a' && b <= 'z' {
				l++
			} else if b >= 'A' && b <= 'Z' {
				u++
			} else if b >= '0' && b <= '9' {
				n++
			} else if b == ' ' {
				sp++
			}
		}

		fmt.Fprintln(writer, l, u, n, sp)
	}
	writer.Flush()
}
