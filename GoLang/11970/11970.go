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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, c, d int
	fmt.Fscanln(reader, &a, &b)
	fmt.Fscanln(reader, &c, &d)

	if c >= b || d <= a {
		fmt.Println(b - a + d - c)
	} else {
		fmt.Println(max(b, d) - min(a, c))
	}
}
