package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func removeBUG(s string) string {
	for strings.Contains(s, "BUG") {
		s = strings.ReplaceAll(s, "BUG", "")
	}
	return s
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

loop:
	for {
		s, err := reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")
		switch err {
		case nil:
		case io.EOF:
			break loop
		}

		fmt.Fprintln(writer, removeBUG(s))
	}

	writer.Flush()
}
