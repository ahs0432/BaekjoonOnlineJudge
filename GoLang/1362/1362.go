package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var rip bool
	var s string
	var o, w, t int
	for i := 1; ; i++ {
		fmt.Fscanln(reader, &o, &w)

		if o == 0 && w == 0 {
			break
		}

		rip = false
		for {
			fmt.Fscanln(reader, &s, &t)
			if s == "#" && t == 0 {
				break
			} else if s == "F" {
				w += t
			} else if s == "E" {
				w -= t
			}

			if w <= 0 {
				rip = true
			}
		}

		if w <= 0 || rip {
			fmt.Fprintln(writer, i, "RIP")
		} else if w > o/2 && w < o*2 {
			fmt.Fprintln(writer, i, ":-)")
		} else {
			fmt.Fprintln(writer, i, ":-(")
		}
	}
	writer.Flush()
}
