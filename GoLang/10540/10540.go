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

	var a, b int
	maxX, minX := -2147483647, 2147483647
	maxY, minY := -2147483647, 2147483647
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b)
		maxX = max(maxX, a)
		minX = min(minX, a)
		maxY = max(maxY, b)
		minY = min(minY, b)
	}
	tmp := max(maxX-minX, maxY-minY)
	fmt.Println(tmp * tmp)
}
