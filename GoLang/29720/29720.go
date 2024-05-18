package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m, k int
	fmt.Fscanln(reader, &n, &m, &k)

	fmt.Println(max(0, n-m*k), max(0, n-m*(k-1)-1))
}
