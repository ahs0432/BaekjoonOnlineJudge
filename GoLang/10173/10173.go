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

	for {
		s, _ = reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")
		if s == "EOI" {
			break
		}

		s = strings.ToLower(s)
		if strings.Contains(s, "nemo") {
			fmt.Fprintln(writer, "Found")
		} else {
			fmt.Fprintln(writer, "Missing")
		}
	}
	writer.Flush()
}
