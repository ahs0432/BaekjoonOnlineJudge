package main

import (
	"bufio"
	"fmt"
	"os"
)

func insert(array []int, index int, i int) []int {
	result := append(array, i)
	copy(result[index+1:], result[index:])
	result[index] = i
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, tmp int
	fmt.Fscanln(reader, &n)

	table := []int{}

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)
		table = insert(table, tmp, i+1)
	}

	for i := n - 1; i >= 0; i-- {
		fmt.Fprint(writer, table[i], " ")
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
