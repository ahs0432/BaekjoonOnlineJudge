package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var x1, y1, x2, y2, x3, y3, x4, y4 int

	fmt.Fscanln(reader, &x1, &y1, &x2, &y2)
	fmt.Fscanln(reader, &x3, &y3, &x4, &y4)

	x := max(x2, x4) - min(x1, x3)
	y := max(y2, y4) - min(y1, y3)
	fmt.Println(max(x, y) * max(x, y))
}
