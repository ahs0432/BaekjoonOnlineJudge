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

	s, err := reader.ReadString('\n')

	if nil != err {
		return
	}

	s = strings.TrimSuffix(s, "\n")

	for _, r := range s {
		if r >= 65 && r <= 90 {
			r = rune((int(r)-52)%26 + 65)
		} else if r >= 97 && r <= 122 {
			r = rune((int(r)-84)%26 + 97)
		}
		fmt.Fprint(writer, string(r))
	}
	writer.Flush()
}
