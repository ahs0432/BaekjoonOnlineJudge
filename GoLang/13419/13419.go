package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var s string
	var s1, s2 string
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s)

		s1, s2 = "", ""

		if len(s)%2 == 0 {
			for j := 0; j < len(s); j++ {
				if j%2 == 0 {
					s1 += string(s[j])
				} else {
					s2 += string(s[j])
				}
			}
		} else {
			for j := 0; j < len(s); j++ {
				if j%2 == 0 {
					s1 += string(s[j])
				} else {
					s2 += string(s[j])
				}
			}

			for j := 0; j < len(s); j++ {
				if j%2 == 0 {
					s2 += string(s[j])
				} else {
					s1 += string(s[j])
				}
			}
		}

		fmt.Fprintln(writer, s1)
		fmt.Fprintln(writer, s2)
	}
	writer.Flush()
}
