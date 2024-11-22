package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var a, b int
	var s string
	var check bool
	for i := 1; ; i++ {
		fmt.Fscanln(reader, &a, &s, &b)
		if s == "E" {
			break
		}

		check = false
		if s == ">" {
			if a > b {
				check = true
			}
		} else if s == ">=" {
			if a >= b {
				check = true
			}
		} else if s == "<" {
			if a < b {
				check = true
			}
		} else if s == "<=" {
			if a <= b {
				check = true
			}
		} else if s == "==" {
			if a == b {
				check = true
			}
		} else if s == "!=" {
			if a != b {
				check = true
			}
		}

		if check {
			fmt.Fprintf(writer, "Case %d: true\n", i)
		} else {
			fmt.Fprintf(writer, "Case %d: false\n", i)
		}
	}
	writer.Flush()
}
