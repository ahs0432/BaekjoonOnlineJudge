package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var x, y, n int
	fmt.Fscanln(reader, &x, &y, &n)

	var check bool
	for i := 1; i <= n; i++ {
		check = false
		if i%x == 0 {
			fmt.Fprint(writer, "Fizz")
			check = true
		}
		if i%y == 0 {
			fmt.Fprint(writer, "Buzz")
			check = true
		}

		if check {
			fmt.Fprintln(writer)
		} else {
			fmt.Fprintln(writer, i)
		}
	}
	writer.Flush()
}
