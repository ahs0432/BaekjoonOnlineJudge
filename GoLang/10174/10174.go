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

	var n int
	fmt.Fscanln(reader, &n)

	var s string
	var t []byte
	var check bool = true
	for i := 0; i < n; i++ {
		s, _ = reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")

		t = []byte{}
		for _, c := range s {
			if c >= 'A' && c <= 'Z' {
				c += 32
			}
			t = append(t, byte(c))
		}

		check = true
		for i := 0; i < len(t)/2; i++ {
			if t[i] != t[len(t)-1-i] {
				check = false
				break
			}
		}

		if check {
			fmt.Fprintln(writer, "Yes")
		} else {
			fmt.Fprintln(writer, "No")
		}
	}
	writer.Flush()
}
