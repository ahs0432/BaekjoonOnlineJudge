package main

import (
	"bufio"
	"fmt"
	"os"
)

func countZero(n, m int) int {
	count := 0
	for i := n; i <= m; i++ {
		if i == 0 {
			count++
		}
		j := i

		for j > 0 {
			if j%10 == 0 {
				count++
			}
			j /= 10
		}
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t, n, m int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n, &m)
		fmt.Fprintln(writer, countZero(n, m))
	}
	writer.Flush()
}
