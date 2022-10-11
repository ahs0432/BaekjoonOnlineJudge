package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var num int
	fmt.Fscanln(reader, &num)

	table := make([]int, 10001)

	var n int
	for i := 0; i < num; i++ {
		fmt.Fscanln(reader, &n)
		table[n]++
	}

	for i := 1; i < len(table); i++ {
		for j := 0; j < table[i]; j++ {
			fmt.Fprintln(writer, i)
		}
	}
}
