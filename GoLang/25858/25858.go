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
	var m int
	fmt.Fscanln(reader, &n, &m)

	sum := 0
	t := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &t[i])
		sum += t[i]
	}

	m /= sum
	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, m*t[i])
	}
	writer.Flush()
}
