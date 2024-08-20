package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	sticks := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &sticks[i])
	}

	count := 0
	maxH := 0
	for i := n - 1; i >= 0; i-- {
		if sticks[i] > maxH {
			count++
			maxH = sticks[i]
		}
	}
	fmt.Println(count)
}
