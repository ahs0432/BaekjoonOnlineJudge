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

	k, _ := reader.ReadString('\n')
	k = strings.TrimSuffix(k, "\n")

	var m int
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			fmt.Fprint(writer, " ")
			continue
		}

		m = int(s[i]-'a') - int((k[i%len(k)] - 'a'))

		if m < 1 {
			fmt.Fprint(writer, string('z'+m))
		} else {
			fmt.Fprint(writer, string('a'+m-1))
		}
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
