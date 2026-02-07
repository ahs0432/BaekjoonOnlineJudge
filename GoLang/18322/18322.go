package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	table := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &table[i])
	}

	count := 0
	for i := 0; i < n; i++ {
		if count+len(table[i]) > k {
			fmt.Fprintln(writer)
			fmt.Fprintf(writer, "%s", table[i])
			count = len(table[i])
		} else {
			if count > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprintf(writer, "%s", table[i])
			count += len(table[i])
		}
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
