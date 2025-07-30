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

	var str string
	var table []string
	for i := 1; i <= n; i++ {
		str, _ = reader.ReadString('\n')
		str = strings.TrimSuffix(str, "\n")
		table = strings.Split(str, " ")

		fmt.Fprintf(writer, "Case #%d: ", i)
		for j := len(table) - 1; j >= 0; j-- {
			fmt.Fprint(writer, table[j], " ")
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
