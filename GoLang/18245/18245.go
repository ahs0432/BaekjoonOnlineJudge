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
	for i := 2; ; i++ {
		s, _ = reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")

		if s == "Was it a cat I saw?" {
			break
		}

		for j := 0; j < len(s); j += i {
			fmt.Fprint(writer, string(s[j]))
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
