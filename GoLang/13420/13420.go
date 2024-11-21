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

	var a, b, ans int
	var s, e string
	var check bool = false
	for i := 0; i < n; i++ {
		check = false
		fmt.Fscanln(reader, &a, &s, &b, &e, &ans)

		if s == "+" {
			if a+b == ans {
				check = true
			}
		} else if s == "-" {
			if a-b == ans {
				check = true
			}
		} else if s == "*" {
			if a*b == ans {
				check = true
			}
		} else if s == "/" {
			if a/b == ans {
				check = true
			}
		}

		if check {
			fmt.Fprintln(writer, "correct")
		} else {
			fmt.Fprintln(writer, "wrong answer")
		}
	}
	writer.Flush()
}
