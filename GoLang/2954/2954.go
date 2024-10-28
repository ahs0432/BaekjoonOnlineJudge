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

	s, _ := reader.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")

	for i := 0; i < len(s); i++ {
		fmt.Fprint(writer, string(s[i]))
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			i += 2
		}
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
