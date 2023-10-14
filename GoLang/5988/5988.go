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
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &s)
		a := int(byte(s[len(s)-1]))
		if a%2 == 0 {
			fmt.Fprintln(writer, "even")
		} else {
			fmt.Fprintln(writer, "odd")
		}
	}

	writer.Flush()
}
