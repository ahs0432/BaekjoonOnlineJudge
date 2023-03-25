package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var x, y int
	fmt.Fscanln(reader, &y, &x)

	table := make([]string, y)

	for i := 0; i < y; i++ {
		fmt.Fscanln(reader, &table[i])
	}

	for i := 0; i < y; i++ {
		for j := x - 1; j >= 0; j-- {
			if j == 0 {
				fmt.Fprintln(writer, string(table[i][j]))
			} else {
				fmt.Fprint(writer, string(table[i][j]))
			}
		}
	}

	writer.Flush()
}
