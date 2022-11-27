package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	table := make([]int, n)

	for i := 0; i < n; i++ {
		table[i] = i + 1
	}

	for len(table) != 1 {
		table = table[1:]
		table = append(table, table[0])
		table = table[1:]
	}

	fmt.Println(table[0])
}
