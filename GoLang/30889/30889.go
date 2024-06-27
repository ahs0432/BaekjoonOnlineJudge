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

	table := make([][]string, 10)

	for i := 0; i < 10; i++ {
		table[i] = make([]string, 20)
		for j := 0; j < 20; j++ {
			table[i][j] = "."
		}
	}

	var tmp string
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)

		i := tmp[0] - 65
		j := tmp[1] - 48

		if len(tmp) == 3 {
			j *= 10
			j += tmp[2] - 48
		}

		table[i][j-1] = "o"
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			fmt.Fprint(writer, table[i][j])
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
