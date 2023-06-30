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

	var a, x, b, y, t int
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &x)
	fmt.Fscanln(reader, &b)
	fmt.Fscanln(reader, &y)
	fmt.Fscanln(reader, &t)

	fmt.Println(a+max(t-30, 0)*x*21, b+max(t-45, 0)*y*21)
}
