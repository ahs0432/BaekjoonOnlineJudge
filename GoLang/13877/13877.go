package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var t, o8, o10, o16 int64
	var s string
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &t, &s)
		o8, _ = strconv.ParseInt(s, 8, 64)
		o10, _ = strconv.ParseInt(s, 10, 64)
		o16, _ = strconv.ParseInt(s, 16, 64)

		fmt.Fprintln(writer, t, o8, o10, o16)
	}

	writer.Flush()
}
