package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	table := make([]int, 80001)
	for i := 3; i <= 80000; i++ {
		table[i] = table[i-1]
		if i%3 == 0 || i%7 == 0 {
			table[i] += i
		}
	}

	var t int
	fmt.Fscanln(reader, &t)

	var n int
	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &n)
		fmt.Fprintln(writer, table[n])
	}
	writer.Flush()
}
