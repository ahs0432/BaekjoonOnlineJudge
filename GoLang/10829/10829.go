package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int64
	fmt.Fscanln(reader, &n)

	s := ""
	for n != 0 {
		s = string(byte(n%2+48)) + s
		n /= 2
	}
	fmt.Fprintln(writer, s)
	writer.Flush()
}
