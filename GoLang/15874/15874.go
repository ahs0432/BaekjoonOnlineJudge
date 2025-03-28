package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var k, sl int
	fmt.Fscanln(reader, &k, &sl)

	var s string
	s, _ = reader.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")

	k %= 26

	var b []byte
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			b = append(b, s[i]+byte(k))

			if b[i] > 'z' {
				b[i] = 'a' + (b[i]%'z' - 1)
			}
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			b = append(b, s[i]+byte(k))

			if b[i] > 'Z' {
				b[i] = 'A' + (b[i]%'Z' - 1)
			}
		} else {
			b = append(b, s[i])
		}
	}
	fmt.Println(string(b))
}
