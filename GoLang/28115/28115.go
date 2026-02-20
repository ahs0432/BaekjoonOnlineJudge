package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(n int, table []int) bool {
	for i := 2; i < n; i++ {
		if table[i]-table[i-1] != table[1]-table[0] {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	table := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &table[i])
	}

	if !check(n, table) {
		fmt.Fprintln(writer, "NO")
	} else {
		fmt.Fprintln(writer, "YES")
		for i := 0; i < n; i++ {
			fmt.Fprint(writer, table[i], " ")
		}
		fmt.Fprintln(writer)
		for i := 0; i < n; i++ {
			fmt.Fprint(writer, "0 ")
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
