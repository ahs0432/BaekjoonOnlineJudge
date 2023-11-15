package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(s string) string {
	tmp := ""
	for _, c := range s {
		tmp = string(c) + tmp
	}
	return tmp
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	table := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &table[i])
	}
	var t int
	fmt.Fscanln(reader, &t)

	for i := 0; i < n; i++ {
		if t == 1 {
			fmt.Fprintln(writer, table[i])
		} else if t == 2 {
			fmt.Fprintln(writer, reverse(table[i]))
		} else if t == 3 {
			fmt.Fprintln(writer, table[n-1-i])
		}
	}

	writer.Flush()
}
