package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b, x, y int
	fmt.Fscanln(reader, &a, &b, &x, &y)
	fmt.Println(min(abs(a-b), min(abs(a-x)+abs(b-y), abs(a-y)+abs(b-x))))
}
