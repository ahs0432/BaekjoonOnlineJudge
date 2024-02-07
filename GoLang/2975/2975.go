package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var b, a int
	var o string

	r := 0
	for {
		fmt.Fscanln(reader, &b, &o, &a)

		if b == 0 && o == "W" && a == 0 {
			break
		}

		r = 0
		if o == "W" {
			r = b - a
		} else {
			r = b + a
		}

		if r < -200 {
			fmt.Fprintln(writer, "Not allowed")
		} else {
			fmt.Fprintln(writer, r)
		}
	}
	writer.Flush()
}
