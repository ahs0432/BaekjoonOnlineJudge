package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var s string
	var count int
	for {
		count = 0
		fmt.Fscanln(reader, &s)
		if s == "#" {
			break
		}

		for i := 0; i < len(s)-1; i++ {
			if s[i] == '1' {
				count++
			}
			fmt.Fprint(writer, string(s[i]))
		}

		if s[len(s)-1] == 'e' {
			if count%2 == 0 {
				fmt.Fprintln(writer, "0")
			} else {
				fmt.Fprintln(writer, "1")
			}
		} else {
			if count%2 == 1 {
				fmt.Fprintln(writer, "0")
			} else {
				fmt.Fprintln(writer, "1")
			}
		}
	}
	writer.Flush()
}
