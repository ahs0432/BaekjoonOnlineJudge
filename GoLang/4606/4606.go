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
		if s == "#" {
			break
		}

		s = strings.ReplaceAll(s, "%", "%25")
		s = strings.ReplaceAll(s, " ", "%20")
		s = strings.ReplaceAll(s, "!", "%21")
		s = strings.ReplaceAll(s, "$", "%24")
		s = strings.ReplaceAll(s, "(", "%28")
		s = strings.ReplaceAll(s, ")", "%29")
		s = strings.ReplaceAll(s, "*", "%2a")
		fmt.Fprintln(writer, s)
	}
	writer.Flush()
}
