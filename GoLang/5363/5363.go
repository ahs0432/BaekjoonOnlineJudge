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
	var sl []string
	for i := 0; i < n; i++ {
		s, _ = reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")

		sl = strings.Split(s, " ")

		for j := 2; j < len(sl); j++ {
			fmt.Fprint(writer, sl[j], " ")
		}

		for j := 0; j < 2; j++ {
			fmt.Fprint(writer, sl[j], " ")
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
