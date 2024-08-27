package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	table := make([]int, 21)
	for i := 1; i < 21; i++ {
		table[i] = i
	}

	var a, b int
	for i := 0; i < 10; i++ {
		fmt.Fscanln(reader, &a, &b)

		for j := 0; j < (b-a+1)/2; j++ {
			table[a+j], table[b-j] = table[b-j], table[a+j]
		}
	}

	for i := 1; i < 21; i++ {
		fmt.Fprint(writer, table[i], " ")
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
