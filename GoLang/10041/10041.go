package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var w, h, n int
	fmt.Fscanln(reader, &w, &h, &n)

	p := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &p[i][0], &p[i][1])
	}

	answer := 0
	for i := 0; i < n-1; i++ {
		x1, y1 := p[i][0], p[i][1]
		x2, y2 := p[i+1][0], p[i+1][1]

		if x1 > x2 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}

		if y1 > y2 {
			answer += abs(x1-x2) + abs(y1-y2)
		} else {
			answer += max(abs(x1-x2), abs(y1-y2))
		}
	}

	fmt.Println(answer)
}
