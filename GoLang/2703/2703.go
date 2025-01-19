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

	var t int
	fmt.Fscanln(reader, &t)

	var a, b string
	var now []byte
	for i := 0; i < t; i++ {
		a, _ = reader.ReadString('\n')
		a = strings.TrimSuffix(a, "\n")
		fmt.Fscanln(reader, &b)

		now = []byte{}
		for _, c := range a {
			if c == ' ' {
				now = append(now, byte(c))
				continue
			}
			now = append(now, b[c-65])
		}
		fmt.Fprintln(writer, string(now))
	}
	writer.Flush()
}
