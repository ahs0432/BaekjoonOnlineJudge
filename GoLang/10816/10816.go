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

	table := map[int]int{}

	for i := 0; i < n; i++ {
		var tmp int
		if i == n-1 {
			fmt.Fscanln(reader, &tmp)
		} else {
			fmt.Fscan(reader, &tmp)
		}
		table[tmp]++
	}

	var m int
	fmt.Fscanln(reader, &m)

	for i := 0; i < m; i++ {
		var tmp int
		if i == m-1 {
			fmt.Fscanln(reader, &tmp)
		} else {
			fmt.Fscan(reader, &tmp)
		}
		fmt.Fprint(writer, table[tmp], " ")
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
