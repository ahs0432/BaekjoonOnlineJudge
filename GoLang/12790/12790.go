package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	t := make([]int, 8)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &t[0], &t[1], &t[2], &t[3], &t[4], &t[5], &t[6], &t[7])

		for j := 0; j < 4; j++ {
			t[j] += t[j+4]
		}

		fmt.Fprintln(writer, max(t[0], 1)+max(t[1], 1)*5+max(t[2], 0)*2+t[3]*2)
	}
	writer.Flush()
}
