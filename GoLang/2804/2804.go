package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var s1, s2 string
	fmt.Fscanln(reader, &s1, &s2)

	c1, c2 := -1, -1
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				c1, c2 = i, j
				break
			}
		}

		if c1 != -1 {
			break
		}
	}

	for i := 0; i < len(s2); i++ {
		if i == c2 {
			fmt.Fprintln(writer, s1)
		} else {
			for j := 0; j < c1; j++ {
				fmt.Fprint(writer, ".")
			}
			fmt.Fprint(writer, string(s2[i]))
			for j := 0; j < len(s1)-c1-1; j++ {
				fmt.Fprint(writer, ".")
			}
			fmt.Fprintln(writer)
		}
	}
	writer.Flush()
}
