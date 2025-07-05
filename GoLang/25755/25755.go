package main

import (
	"bufio"
	"fmt"
	"os"
)

func changeNum(i int) int {
	if i == 2 {
		return 5
	} else if i == 5 {
		return 2
	} else if i == 1 || i == 8 {
		return i
	}

	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var w string
	var n int
	fmt.Fscanln(reader, &w, &n)

	table := make([][]int, n)

	for i := 0; i < n; i++ {
		table[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if w == "L" || w == "R" {
				fmt.Fscan(reader, &table[i][n-j-1])
				table[i][n-j-1] = changeNum(table[i][n-j-1])
			} else {
				fmt.Fscan(reader, &table[n-i-1][j])
				table[n-i-1][j] = changeNum(table[n-i-1][j])
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if table[i][j] == 0 {
				fmt.Fprint(writer, "? ")
			} else {
				fmt.Fprintf(writer, "%d ", table[i][j])
			}
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
