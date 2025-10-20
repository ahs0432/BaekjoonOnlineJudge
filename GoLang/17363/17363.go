package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	table := make([][]byte, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &table[i])
	}

	for i := m - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			switch table[j][i] {
			case '-':
				table[j][i] = '|'
			case '|':
				table[j][i] = '-'
			case '/':
				table[j][i] = '\\'
			case '\\':
				table[j][i] = '/'
			case '^':
				table[j][i] = '<'
			case '<':
				table[j][i] = 'v'
			case 'v':
				table[j][i] = '>'
			case '>':
				table[j][i] = '^'
			}
			fmt.Fprint(writer, string(table[j][i]))
		}
		fmt.Fprintln(writer)
	}

	writer.Flush()
}
