package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var s string
	for i := 1; i <= n; i++ {
		s, _ = reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")
		ss := strings.Split(s, " ")

		fmt.Fprintf(writer, "Case #%d: ", i)
		for i := len(ss) - 1; i >= 0; i-- {
			fmt.Fprint(writer, ss[i], " ")
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
